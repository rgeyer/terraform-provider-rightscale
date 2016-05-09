package main

import (
	"io"
  "os"
  "fmt"
  "os/exec"
	"strings"
  "regexp"
	"text/template"

	"gopkg.in/rightscale/rsc.v5/metadata"
  "gopkg.in/rightscale/rsc.v5/cm15"
  "bitbucket.org/pkg/inflect"
)

// TerraformWriter struct exposes methods to generate the go API client command line tool
type TerraformWriter struct {
	headerTmpl   *template.Template
	resourceTmpl *template.Template
}

// NewTerraformWriter creates a new writer that generates metadata data structures.
func NewTerraformWriter() (*TerraformWriter, error) {
	funcMap := template.FuncMap{
    "underscore":       inflect.Underscore,
    "camelize":         inflect.Camelize,
		"join":            	strings.Join,
		"isCreateOrUpdate":	stringIsCreateOrUpdate,
		"terraformSchemaAttributes": terraformSchemaAttributes,
	}
	headerT, err := template.New("header-terraform").Funcs(funcMap).Parse(headerTerraformTmpl)
	if err != nil {
		return nil, err
	}
	resourceT, err := template.New("resource-terraform").Funcs(funcMap).Parse(resourceTerraformTmpl)
	if err != nil {
		return nil, err
	}
	return &TerraformWriter{
		headerTmpl:   headerT,
		resourceTmpl: resourceT,
	}, nil
}

// WriteHeader writes the generic header text.
func (c *TerraformWriter) WriteHeader(pkg string, w io.Writer) error {
	return c.headerTmpl.Execute(w, pkg)
}

// WriteMetadata writes the data structures that describe the API resources and actions.
func (c *TerraformWriter) WriteMetadata(d map[string]*metadata.Resource, w io.Writer) error {
	return c.resourceTmpl.Execute(w, d)
}

func stringIsCreateOrUpdate(s string) bool {
	l := []string{"create","update"}
  for _, b := range l {
      if b == s {
          return true
      }
  }
  return false
}

type TerraformSchemaAttribute struct {
	Name string
	SchemaType   string
	Required bool
	Computed bool
	ForceNew bool
  Elem string
  Children map[string]*TerraformSchemaAttribute
	Debug string
}

func (t *TerraformSchemaAttribute) ShouldPrintChildren() bool {
	return t.Children != nil && len(t.Children) > 0
}

func (t *TerraformSchemaAttribute) RequiredProp() string {
	if t.Required && !t.Computed {
		return "Required: true"
	} else {
		return "Optional: true"
	}
}

func (t *TerraformSchemaAttribute) SetSchemaTypes(schema_type string) {
  list_regex := regexp.MustCompile(`(\[\])(.*)`)
  t.SchemaType = "schema.TypeString"
  switch schema_type {
  case "string", "*RubyTime", "file":
    t.SchemaType = "schema.TypeString"
  case "bool":
    t.SchemaType = "schema.TypeBool"
  case "int":
    t.SchemaType = "schema.TypeInt"
  case "[]string":
    t.SchemaType = "schema.TypeList"
    t.Elem = "&schema.Schema{Type: schema.TypeString}"
  case "[]map[string]string","[]map[string]interface{}":
    t.SchemaType = "schema.TypeList"
    t.Elem = "&schema.Schema{Type: schema.TypeMap}"
  case "map[string]string","map[string]interface{}","map":
    t.SchemaType = "schema.TypeMap"
  default:
    matches := list_regex.FindStringSubmatch(schema_type)
    if matches != nil {
      if matches[1] == `[]` {
        t.SchemaType = "schema.TypeList, //"+schema_type
        t.Elem = "&schema.Schema{Type: schema.TypeMap}"
      }
    } else {
      // TODO: This is hacky and broken, but useful for feedback/debugging
      t.SchemaType = "nil, //"+schema_type
    }
  }
}

func terraformSchemaAttributes(r *metadata.Resource) []*TerraformSchemaAttribute {
  par_attr := make(map[string]*TerraformSchemaAttribute)
	top_level_attr := []*TerraformSchemaAttribute{}
	values := []*TerraformSchemaAttribute{}

  has_create	:= false
  has_update	:= false
	// has_read		:= false

	for _, action := range r.Actions {
		if action.Name != "create" && action.Name != "update" {
      continue
    }
    // fmt.Println("Interrogating action "+action.Name)
    for _, cmdFlag := range action.CommandFlags {
			var temp_attr *TerraformSchemaAttribute
      // fmt.Println(fmt.Sprintf("%s - %s", cmdFlag.Name, cmdFlag.Type))
      param_re := regexp.MustCompile(`([a-z_]*)\[?([a-z_]*)\]?`)
      submatches := param_re.FindAllStringSubmatch(cmdFlag.Name, -1)

      // It always should be, but let's check to be sure
      if len(submatches) >= 1 {
        parent_map := par_attr
				operating_at_root := true
				// last_child_key := ""
        for subMatchIdx, subMatch := range submatches {
          this_key := subMatch[1]
          if subMatch[1] == "" {
            this_key = subMatch[2]
          }
          if val, ok := parent_map[this_key]; ok {
            temp_attr = val
            // fmt.Println(fmt.Sprintf("Updating field %s in map at key %s", val.Name, this_key))
						temp_attr.Debug = fmt.Sprintf("%s\n// Discovered at key %s of parent map and updating -- Submatch: %q at idx %d -- Operating at root: %t", temp_attr.Debug, this_key, subMatch, subMatchIdx, operating_at_root)
          } else {
						// fmt.Println(fmt.Sprintf("Creating temp_attr with name %s", this_key))
            temp_attr = &TerraformSchemaAttribute{
              Name: this_key,
            }
						temp_attr.SetSchemaTypes(cmdFlag.Type)
						temp_attr.Debug = fmt.Sprintf("Created for key %s. -- Submatch: %q at idx %d -- cmdFlagName: %s -- Operating at root: %t", this_key, subMatch, subMatchIdx, cmdFlag.Name, operating_at_root)
          }

          if subMatch[2] == "" {
            temp_attr.SetSchemaTypes(cmdFlag.Type)
						temp_attr.Debug = fmt.Sprintf("%s\n// Matched at subMatch[2] == '' -- Operating at root: %t", temp_attr.Debug, operating_at_root)
          } else if subMatchIdx  == (len(submatches)-1) {
						temp_attr.Debug = fmt.Sprintf("%s\n// Matched at subMatchIdx == last -- Operating at root: %t", temp_attr.Debug, operating_at_root)
					} else {
            child := subMatch[2]
						// fmt.Println(fmt.Sprintf("Assigning child %s to parent %s using submatch %q", child, this_key, subMatch))
						temp_attr.SchemaType = "schema.TypeSet"
						if temp_attr.Children == nil {
							temp_attr.Debug = fmt.Sprintf("%s\n// My children were nil and are being initialized", temp_attr.Debug)
							temp_attr.Children = make(map[string]*TerraformSchemaAttribute)
						}
						if _, ok := temp_attr.Children[child]; !ok {
							// fmt.Println(fmt.Sprintf("Creating temp_attr with name %s at child of %s", child, this_key))
							temp_attr.Children[child] = &TerraformSchemaAttribute{
								Name: child,
								Debug: fmt.Sprintf("Child at parent[%s][%s] -- Submatch: %q at idx %d -- cmdFlagName: %s", this_key, child, subMatch, subMatchIdx, cmdFlag.Name),
							}
							temp_attr.Children[child].SetSchemaTypes(cmdFlag.Type)
							temp_attr.Debug = fmt.Sprintf("Created as a child node when an existing child of the same key was not found. Parent was %s -- Submatch: %q at idx %d -- Operating at root: %t", this_key, subMatch, subMatchIdx, operating_at_root)
						}
						if operating_at_root {
							temp_attr.Required = true
							temp_attr.Debug = fmt.Sprintf("%s\n// This is a top level attribute.", temp_attr.Debug)
							top_level_attr = append(top_level_attr, temp_attr)
							operating_at_root = false

							// fmt.Println(fmt.Sprintf("Only once, assigning child: %s, this_key: %s", child, this_key))

							if temp_attr.Children[child].Children == nil {
								fmt.Println("This is an infinite loop?")
								current_child_keys := []string{}
								for key, _ := range temp_attr.Children[child].Children {
									current_child_keys = append(current_child_keys, key)
								}
								temp_attr.Children[child].Debug = fmt.Sprintf("%s\n// My children are being initialized, they were %q", temp_attr.Children[child].Debug, current_child_keys)
								temp_attr.Children[child].Children = make(map[string]*TerraformSchemaAttribute)
							}
							fmt.Println("Finished allocating children")
							temp_attr.Children[child].SchemaType = "schema.TypeSet"
							parent_map[this_key] = temp_attr
							temp_attr.Children[child].Debug = fmt.Sprintf("%s\n// (At Root) My children became the parent map. -- Submatch: %q -- CmdFlag: %s", temp_attr.Children[child].Debug, subMatch, cmdFlag.Name)
							parent_map = temp_attr.Children[child].Children
							temp_attr = nil
							fmt.Println("Finished setting temp_attr to nil")
						} else {
							temp_attr.Debug = fmt.Sprintf("%s\n// (Below Root) My children became the parent map. -- Submatch: %q -- CmdFlag: %s", temp_attr.Debug, subMatch, cmdFlag.Name)
							parent_map = temp_attr.Children
							temp_attr = nil
						}
          }

          if temp_attr != nil {
						if temp_attr.Name == "" {
							fmt.Println(fmt.Sprintf("The attribute name was blank for %s", cmdFlag.Name))
						} else {
							if _, ok := parent_map[this_key]; ok {
								fmt.Println(fmt.Sprintf("Submatch IDX is %d, and we'd overwrite %s", subMatchIdx, this_key))
							} else {
								fmt.Println(fmt.Sprintf("Submatch IDX is %d, and assigning %s to parent map at %s", subMatchIdx, temp_attr.Name, this_key))
							}
							temp_attr.Debug = fmt.Sprintf("%s\n// Assigned to parent map at %s because temp_attr wasn't null, which it never will be", temp_attr.Debug, this_key)
							parent_map[this_key] = temp_attr
						}
          } else {
            fmt.Println(fmt.Sprintf("Nothing to assign for key %s", this_key))
          }

          // if temp_attr != nil && temp_attr.Children != nil {
					// 	fmt.Println(fmt.Sprintf("Parent map is now updated to the node with the key %s, which has children", temp_attr.Name))
          //   parent_map = temp_attr.Children
          // }
        }
      }
      fmt.Println("\n\n")
			if action.Name == "show" {
				// has_read = true
			}
	    if action.Name == "update" {
				has_update = true
				if temp_attr != nil && cmdFlag.Mandatory {
					temp_attr.Required = true
				}
	    }
	    if action.Name == "create" {
				has_create = true
				if temp_attr != nil && cmdFlag.Mandatory {
					temp_attr.Required = true
				}
	    }
    }
	}

	for _, attribute := range r.Attributes {
		if len(top_level_attr) > 0 {
			if _, ok := top_level_attr[0].Children[attribute.Name]; ok {
				// A property at the root has a child property with this same name.
				// This is meant to match stuff like deployment[name] and name, which
				// are actually the same thing.
				continue
			}
		}
		var attr *TerraformSchemaAttribute
		if val, ok := par_attr[attribute.Name]; ok {
			attr = val
		} else {
	    attr = &TerraformSchemaAttribute{
	      Name: attribute.Name,
	      Computed: true,
	      Required: false,
	    }
		}
    attr.SetSchemaTypes(attribute.FieldType)
    par_attr[attribute.Name] = attr
  }

	has_children := 0
	for _, val := range par_attr {
		if val.Children != nil && len(val.Children) > 0 {
			has_children = has_children + 1
		}
    // TODO: This is a bit bold, we're just stripping anything that's not
    // handled. Could (maybe should?) use reflection to more concretely
    // define these inputs, though I think most of these are calculated values.
    if !strings.HasPrefix(val.SchemaType,"nil") {
      if has_create && !has_update && !val.Computed {
        val.ForceNew = true
      }
  		values = append(values, val)
    }
	}
	fmt.Println(fmt.Sprintf("Got %d properties with children", has_children))
  fmt.Println(fmt.Sprintf("----------------The full map for this looks like %q\n\n\n", par_attr))
	return values
}

const headerTerraformTmpl = `//************************************************************************//
//                     rsc - RightScale API command line tool
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package main

import (
	"fmt"
  "log"
	"io/ioutil"
	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/rightscale/rsc.v5/cm15"
	"gopkg.in/rightscale/rsc.v5/rsapi"
)
`

const resourceTerraformTmpl = `{{define "schemaObject"}}// DEBUG INFO {{.Debug}}
"{{.Name}}": &schema.Schema{
		Type: {{.SchemaType}},
		{{.RequiredProp}},{{if .Computed}}
		Computed: true,{{end}}{{if .ShouldPrintChildren}}
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				{{range .Children}}{{template "schemaObject" .}},
				{{end}}
			},
		},{{else}}{{if .Elem}}
		Elem: {{.Elem}},{{end}}{{end}}{{if .ForceNew}}
		ForceNew: true,{{end}}
	}{{end}}{{range .}}{{$resource := .}}
func resourceRightScale{{$resource.Name}}() *schema.Resource {
	return &schema.Resource{ {{range .Actions}}{{$action := .}}
		// ACTION: {{$action.Name}}
		//
		// Path Param Names {{join .PathParamNames ","}}
		// Query Param Names {{join .QueryParamNames ","}}
		// Payload Param Names {{join .PayloadParamNames ","}}
		{{range .PathPatterns}}{{if isCreateOrUpdate $action.Name}}// Pattern:{{.Pattern}} Vars:{{join .Variables ","}}{{end}}
		{{end}}
	{{if eq .Name "create"}}Create: resourceRightScale{{$resource.Name}}Create,
{{end}}{{if eq .Name "show"}}Read: resourceRightScale{{$resource.Name}}Read,
{{end}}{{if eq .Name "update"}}Update: resourceRightScale{{$resource.Name}}Update,
{{end}}{{if eq .Name "destroy"}}Delete: resourceRightScale{{$resource.Name}}Delete,
{{end}}{{end}}

		Schema: map[string]*schema.Schema{ {{range terraformSchemaAttributes .}}
			{{template "schemaObject" .}},
		{{end}}},
	}
}

{{range .Actions}}{{$action := .}}
{{if eq .Name "create"}}func resourceRightScale{{$resource.Name}}Create(d *schema.ResourceData, meta interface{}) error {
  client := meta.(*cm15.API)
  {{$firstPath := index $action.PathPatterns 0}}
	params := rsapi.APIParams{}
	req, err := client.BuildRequest("{{$resource.Name}}", "{{$action.Name}}", "{{$firstPath.Pattern}}", params)
	if err != nil {
		message := fmt.Sprintf("{{$resource.Name}} Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("{{$resource.Name}} Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
			message := fmt.Sprintf("{{$resource.Name}} HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
    }
		message := fmt.Sprintf("{{$resource.Name}} Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScale{{$resource.Name}}Read(d, meta)
	return nil
}{{end}}{{if eq .Name "show"}}func resourceRightScale{{$resource.Name}}Read(d *schema.ResourceData, meta interface{}) error {
  return nil
}{{end}}{{if eq .Name "update"}}func resourceRightScale{{$resource.Name}}Update(d *schema.ResourceData, meta interface{}) error {
  return nil
}{{end}}{{if eq .Name "destroy"}}func resourceRightScale{{$resource.Name}}Delete(d *schema.ResourceData, meta interface{}) error {
  return nil
}{{end}}{{end}}
{{end}}

var ResourceList = map[string]*schema.Resource{ {{range .}}{{$resource := .}}
  "rightscale_{{underscore $resource.Name}}": resourceRightScale{{$resource.Name}}(),
{{end}}}
`

func main() {
  f, err := os.Create("terraform_provider_rightscale/codegen_resources.go")
	if err != nil {
    fmt.Println(err.Error())
		return
	}
	c, err := NewTerraformWriter()
	if err != nil {
    fmt.Println(err.Error())
		return
	}
	err = c.WriteHeader("cm15", f)
	if err != nil {
    fmt.Println(err.Error())
		return
	}
	err = c.WriteMetadata(cm15.GenMetadata, f)
	if err != nil {
    fmt.Println(err.Error())
		return
	}
	f.Close()
	o, err := exec.Command("go", "fmt", "terraform_provider_rightscale/codegen_resources.go").CombinedOutput()
	if err != nil {
		fmt.Println("Failed to format generated terraform code:\n", o)
	}
}
