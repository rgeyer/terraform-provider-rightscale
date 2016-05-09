//************************************************************************//
//                     rsc - RightScale API command line tool
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package main

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/rightscale/rsc.v5/cm15"
	"gopkg.in/rightscale/rsc.v5/rsapi"
	"io/ioutil"
	"log"
)

func resourceRightScaleAccount() *schema.Resource {
	return &schema.Resource{
		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleAccountRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"permissions": &schema.Schema{
				Type:     schema.TypeList, //[]Permission,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"products": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleAccountRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleAccountGroup() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleAccountGroupRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleAccountGroupRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleAlert() *schema.Resource {
	return &schema.Resource{
		// ACTION: disable
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: enable
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: quench
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names duration

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleAlertRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"quenched_until": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"triggered_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleAlertRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleAlertSpec() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names alert_spec
		// Pattern:/api/servers/%s/alert_specs Vars:server_id
		// Pattern:/api/server_arrays/%s/alert_specs Vars:server_array_id
		// Pattern:/api/server_templates/%s/alert_specs Vars:server_template_id
		// Pattern:/api/alert_specs Vars:

		Create: resourceRightScaleAlertSpecCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleAlertSpecDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names with_inherited

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleAlertSpecRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names alert_spec
		// Pattern:/api/servers/%s/alert_specs/%s Vars:server_id,id
		// Pattern:/api/server_arrays/%s/alert_specs/%s Vars:server_array_id,id
		// Pattern:/api/server_templates/%s/alert_specs/%s Vars:server_template_id,id
		// Pattern:/api/alert_specs/%s Vars:id

		Update: resourceRightScaleAlertSpecUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"vote_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"vote_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"condition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"variable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key alert_spec. -- Submatch: ["alert_spec[escalation_name]" "alert_spec" "escalation_name"] at idx 0 -- cmdFlagName: alert_spec[escalation_name] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[subject_href]" "alert_spec" "subject_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[description]" "alert_spec" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[vote_type]" "alert_spec" "vote_type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[condition]" "alert_spec" "condition"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[threshold]" "alert_spec" "threshold"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[duration]" "alert_spec" "duration"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[variable]" "alert_spec" "variable"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[vote_tag]" "alert_spec" "vote_tag"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[name]" "alert_spec" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[file]" "alert_spec" "file"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[escalation_name]" "alert_spec" "escalation_name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[description]" "alert_spec" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[condition]" "alert_spec" "condition"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[threshold]" "alert_spec" "threshold"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[vote_type]" "alert_spec" "vote_type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[duration]" "alert_spec" "duration"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[variable]" "alert_spec" "variable"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[vote_tag]" "alert_spec" "vote_tag"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[name]" "alert_spec" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			// Discovered at key alert_spec of parent map and updating -- Submatch: ["alert_spec[file]" "alert_spec" "file"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at alert_spec because temp_attr wasn't null, which it never will be
			"alert_spec": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"escalation_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleAlertSpecCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("AlertSpec", "create", "/api/servers/%s/alert_specs", params)
	if err != nil {
		message := fmt.Sprintf("AlertSpec Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("AlertSpec Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("AlertSpec HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("AlertSpec Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleAlertSpecRead(d, meta)
	return nil
}
func resourceRightScaleAlertSpecDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleAlertSpecRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleAlertSpecUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleAuditEntry() *schema.Resource {
	return &schema.Resource{
		// ACTION: append
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names detail,notify,offset,summary

		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names audit_entry,notify,user_email
		// Pattern:/api/audit_entries Vars:

		Create: resourceRightScaleAuditEntryCreate,

		// ACTION: detail
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names end_date,limit,start_date

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleAuditEntryRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names audit_entry,notify
		// Pattern:/api/audit_entries/%s Vars:id

		Update: resourceRightScaleAuditEntryUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key notify. -- Submatch: ["notify" "notify" ""] at idx 0 -- cmdFlagName: notify -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at notify because temp_attr wasn't null, which it never will be
			// Discovered at key notify of parent map and updating -- Submatch: ["notify" "notify" ""] at idx 0 -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at notify because temp_attr wasn't null, which it never will be
			"notify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"detail_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key audit_entry. -- Submatch: ["audit_entry[auditee_href]" "audit_entry" "auditee_href"] at idx 0 -- cmdFlagName: audit_entry[auditee_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at audit_entry because temp_attr wasn't null, which it never will be
			// Discovered at key audit_entry of parent map and updating -- Submatch: ["audit_entry[summary]" "audit_entry" "summary"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at audit_entry because temp_attr wasn't null, which it never will be
			// Discovered at key audit_entry of parent map and updating -- Submatch: ["audit_entry[detail]" "audit_entry" "detail"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at audit_entry because temp_attr wasn't null, which it never will be
			// Discovered at key audit_entry of parent map and updating -- Submatch: ["audit_entry[summary]" "audit_entry" "summary"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at audit_entry because temp_attr wasn't null, which it never will be
			// Discovered at key audit_entry of parent map and updating -- Submatch: ["audit_entry[offset]" "audit_entry" "offset"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at audit_entry because temp_attr wasn't null, which it never will be
			"audit_entry": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO Created for key user_email. -- Submatch: ["user_email" "user_email" ""] at idx 0 -- cmdFlagName: user_email -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at user_email because temp_attr wasn't null, which it never will be
			"user_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRightScaleAuditEntryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("AuditEntry", "create", "/api/audit_entries", params)
	if err != nil {
		message := fmt.Sprintf("AuditEntry Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("AuditEntry Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("AuditEntry HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("AuditEntry Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleAuditEntryRead(d, meta)
	return nil
}

func resourceRightScaleAuditEntryRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleAuditEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleBackup() *schema.Resource {
	return &schema.Resource{
		// ACTION: cleanup
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names cloud_href,dailies,keep_last,lineage,monthlies,weeklies,yearlies

		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names backup
		// Pattern:/api/backups Vars:

		Create: resourceRightScaleBackupCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleBackupDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names lineage

		// ACTION: restore
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names backup,instance_href

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleBackupRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names backup
		// Pattern:/api/backups/%s Vars:id

		Update: resourceRightScaleBackupUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"committed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"lineage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"volume_snapshot_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"volume_snapshots": &schema.Schema{
				Type:     schema.TypeList, //[]VolumeSnapshot,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO Created as a child node when an existing child of the same key was not found. Parent was backup -- Submatch: ["backup[volume_attachment_hrefs]" "backup" "volume_attachment_hrefs"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key backup of parent map and updating -- Submatch: ["backup[description]" "backup" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at backup because temp_attr wasn't null, which it never will be
			// Discovered at key backup of parent map and updating -- Submatch: ["backup[from_master]" "backup" "from_master"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at backup because temp_attr wasn't null, which it never will be
			// Discovered at key backup of parent map and updating -- Submatch: ["backup[lineage]" "backup" "lineage"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at backup because temp_attr wasn't null, which it never will be
			// Discovered at key backup of parent map and updating -- Submatch: ["backup[name]" "backup" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at backup because temp_attr wasn't null, which it never will be
			// Discovered at key backup of parent map and updating -- Submatch: ["backup[committed]" "backup" "committed"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at backup because temp_attr wasn't null, which it never will be
			"backup": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// DEBUG INFO Child at parent[backup][volume_attachment_hrefs] -- Submatch: ["backup[volume_attachment_hrefs]" "backup" "volume_attachment_hrefs"] at idx 0 -- cmdFlagName: backup[volume_attachment_hrefs][]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["backup[volume_attachment_hrefs]" "backup" "volume_attachment_hrefs"] -- CmdFlag: backup[volume_attachment_hrefs][]
						"volume_attachment_hrefs": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			// DEBUG INFO
			"completed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"from_master": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleBackupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Backup", "create", "/api/backups", params)
	if err != nil {
		message := fmt.Sprintf("Backup Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Backup Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Backup HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Backup Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleBackupRead(d, meta)
	return nil
}
func resourceRightScaleBackupDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleBackupRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleChildAccount() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names child_account
		// Pattern:/api/child_accounts Vars:

		Create: resourceRightScaleChildAccountCreate,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names child_account
		// Pattern:/api/accounts/%s Vars:id
		// Pattern:/api/child_accounts/%s Vars:id

		Update: resourceRightScaleChildAccountUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key child_account. -- Submatch: ["child_account[cluster_href]" "child_account" "cluster_href"] at idx 0 -- cmdFlagName: child_account[cluster_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at child_account because temp_attr wasn't null, which it never will be
			// Discovered at key child_account of parent map and updating -- Submatch: ["child_account[name]" "child_account" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at child_account because temp_attr wasn't null, which it never will be
			// Discovered at key child_account of parent map and updating -- Submatch: ["child_account[name]" "child_account" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at child_account because temp_attr wasn't null, which it never will be
			"child_account": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceRightScaleChildAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("ChildAccount", "create", "/api/child_accounts", params)
	if err != nil {
		message := fmt.Sprintf("ChildAccount Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("ChildAccount Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("ChildAccount HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("ChildAccount Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleChildAccountRead(d, meta)
	return nil
}

func resourceRightScaleChildAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleCloud() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleCloudRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"capabilities": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"cloud_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleCloudRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleCloudAccount() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names cloud_account
		// Pattern:/api/cloud_accounts Vars:

		Create: resourceRightScaleCloudAccountCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleCloudAccountDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleCloudAccountRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key cloud_account. -- Submatch: ["cloud_account[cloud_href]" "cloud_account" "cloud_href"] at idx 0 -- cmdFlagName: cloud_account[cloud_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at cloud_account because temp_attr wasn't null, which it never will be
			// Discovered at key cloud_account of parent map and updating -- Submatch: ["cloud_account[creds]" "cloud_account" "creds"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at cloud_account because temp_attr wasn't null, which it never will be
			// Discovered at key cloud_account of parent map and updating -- Submatch: ["cloud_account[token]" "cloud_account" "token"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at cloud_account because temp_attr wasn't null, which it never will be
			"cloud_account": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleCloudAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("CloudAccount", "create", "/api/cloud_accounts", params)
	if err != nil {
		message := fmt.Sprintf("CloudAccount Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("CloudAccount Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("CloudAccount HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("CloudAccount Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleCloudAccountRead(d, meta)
	return nil
}
func resourceRightScaleCloudAccountDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleCloudAccountRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleCookbook() *schema.Resource {
	return &schema.Resource{
		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleCookbookDelete,

		// ACTION: follow
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names value

		// ACTION: freeze
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names value

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: obsolete
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names value

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleCookbookRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"metadata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"source_info_summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleCookbookDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleCookbookRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleCookbookAttachment() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names cookbook_attachment
		// Pattern:/api/cookbooks/%s/cookbook_attachments Vars:cookbook_id
		// Pattern:/api/server_templates/%s/cookbook_attachments Vars:server_template_id
		// Pattern:/api/cookbook_attachments Vars:

		Create: resourceRightScaleCookbookAttachmentCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleCookbookAttachmentDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		// ACTION: multi_attach
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names cookbook_attachments

		// ACTION: multi_detach
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names cookbook_attachments

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleCookbookAttachmentRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key cookbook_attachment. -- Submatch: ["cookbook_attachment[server_template_href]" "cookbook_attachment" "server_template_href"] at idx 0 -- cmdFlagName: cookbook_attachment[server_template_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at cookbook_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key cookbook_attachment of parent map and updating -- Submatch: ["cookbook_attachment[cookbook_href]" "cookbook_attachment" "cookbook_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at cookbook_attachment because temp_attr wasn't null, which it never will be
			"cookbook_attachment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"dependency": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleCookbookAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("CookbookAttachment", "create", "/api/cookbooks/%s/cookbook_attachments", params)
	if err != nil {
		message := fmt.Sprintf("CookbookAttachment Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("CookbookAttachment Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("CookbookAttachment HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("CookbookAttachment Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleCookbookAttachmentRead(d, meta)
	return nil
}
func resourceRightScaleCookbookAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleCookbookAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleCredential() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names credential
		// Pattern:/api/credentials Vars:

		Create: resourceRightScaleCredentialCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleCredentialDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleCredentialRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names credential
		// Pattern:/api/credentials/%s Vars:id

		Update: resourceRightScaleCredentialUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key credential. -- Submatch: ["credential[description]" "credential" "description"] at idx 0 -- cmdFlagName: credential[description] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at credential because temp_attr wasn't null, which it never will be
			// Discovered at key credential of parent map and updating -- Submatch: ["credential[value]" "credential" "value"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at credential because temp_attr wasn't null, which it never will be
			// Discovered at key credential of parent map and updating -- Submatch: ["credential[name]" "credential" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at credential because temp_attr wasn't null, which it never will be
			// Discovered at key credential of parent map and updating -- Submatch: ["credential[description]" "credential" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at credential because temp_attr wasn't null, which it never will be
			// Discovered at key credential of parent map and updating -- Submatch: ["credential[value]" "credential" "value"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at credential because temp_attr wasn't null, which it never will be
			// Discovered at key credential of parent map and updating -- Submatch: ["credential[name]" "credential" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at credential because temp_attr wasn't null, which it never will be
			"credential": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleCredentialCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Credential", "create", "/api/credentials", params)
	if err != nil {
		message := fmt.Sprintf("Credential Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Credential Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Credential HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Credential Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleCredentialRead(d, meta)
	return nil
}
func resourceRightScaleCredentialDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleCredentialRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleCredentialUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleDatacenter() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleDatacenterRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleDatacenterRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleDeployment() *schema.Resource {
	return &schema.Resource{
		// ACTION: clone
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names deployment

		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names deployment
		// Pattern:/api/deployments Vars:

		Create: resourceRightScaleDeploymentCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleDeploymentDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: lock
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: servers
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleDeploymentRead,

		// ACTION: unlock
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names deployment
		// Pattern:/api/deployments/%s Vars:id

		Update: resourceRightScaleDeploymentUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"inputs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"locked": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"server_tag_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key deployment. -- Submatch: ["deployment[server_tag_scope]" "deployment" "server_tag_scope"] at idx 0 -- cmdFlagName: deployment[server_tag_scope] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at deployment because temp_attr wasn't null, which it never will be
			// Discovered at key deployment of parent map and updating -- Submatch: ["deployment[description]" "deployment" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at deployment because temp_attr wasn't null, which it never will be
			// Discovered at key deployment of parent map and updating -- Submatch: ["deployment[name]" "deployment" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at deployment because temp_attr wasn't null, which it never will be
			// Discovered at key deployment of parent map and updating -- Submatch: ["deployment[server_tag_scope]" "deployment" "server_tag_scope"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at deployment because temp_attr wasn't null, which it never will be
			// Discovered at key deployment of parent map and updating -- Submatch: ["deployment[description]" "deployment" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at deployment because temp_attr wasn't null, which it never will be
			// Discovered at key deployment of parent map and updating -- Submatch: ["deployment[name]" "deployment" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at deployment because temp_attr wasn't null, which it never will be
			"deployment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceRightScaleDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Deployment", "create", "/api/deployments", params)
	if err != nil {
		message := fmt.Sprintf("Deployment Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Deployment Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Deployment HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Deployment Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleDeploymentRead(d, meta)
	return nil
}
func resourceRightScaleDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleDeploymentUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleHealthCheck() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Schema: map[string]*schema.Schema{},
	}
}

func resourceRightScaleIdentityProvider() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleIdentityProviderRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"discovery_hint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleIdentityProviderRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleImage() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleImageRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"cpu_architecture": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"image_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"root_device_storage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"virtualization_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"os_platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleImageRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleInput() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		// ACTION: multi_update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names inputs

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleInstance() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names api_behavior,instance
		// Pattern:/api/clouds/%s/instances Vars:cloud_id

		Create: resourceRightScaleInstanceCreate,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: launch
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names api_behavior,count,inputs

		// ACTION: lock
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: multi_run_executable
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names ignore_lock,inputs,recipe_name,right_script_href

		// ACTION: multi_terminate
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names terminate_all

		// ACTION: reboot
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: run_executable
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names ignore_lock,inputs,recipe_name,right_script_href

		// ACTION: set_custom_lodgement
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names quantity,timeframe

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleInstanceRead,

		// ACTION: start
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: stop
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: terminate
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: unlock
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names instance
		// Pattern:/api/clouds/%s/instances/%s Vars:cloud_id,id

		Update: resourceRightScaleInstanceUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"private_ip_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			// DEBUG INFO Created as a child node when an existing child of the same key was not found. Parent was instance -- Submatch: ["instance[subnet_hrefs]" "instance" "subnet_hrefs"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[image_href]" "instance" "image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[user_data]" "instance" "user_data"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[name]" "instance" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[associate_public_ip_address]" "instance" "associate_public_ip_address"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[multi_cloud_image_href]" "instance" "multi_cloud_image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[ip_forwarding_enabled]" "instance" "ip_forwarding_enabled"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[server_template_href]" "instance" "server_template_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[security_group_hrefs]" "instance" "security_group_hrefs"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[ramdisk_image_href]" "instance" "ramdisk_image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[instance_type_href]" "instance" "instance_type_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[kernel_image_href]" "instance" "kernel_image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[datacenter_href]" "instance" "datacenter_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[deployment_href]" "instance" "deployment_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[subnet_hrefs]" "instance" "subnet_hrefs"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[ssh_key_href]" "instance" "ssh_key_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[image_href]" "instance" "image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[user_data]" "instance" "user_data"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			// Discovered at key instance of parent map and updating -- Submatch: ["instance[name]" "instance" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at instance because temp_attr wasn't null, which it never will be
			"instance": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// DEBUG INFO Child at parent[instance][cloud_specific_attributes] -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] at idx 0 -- cmdFlagName: instance[cloud_specific_attributes][automatic_instance_store_mapping]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][automatic_instance_store_mapping]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_performance]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_performance]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_type_uid]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_type_uid]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][iam_instance_profile]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][iam_instance_profile]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_size]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_size]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][ebs_optimized]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][ebs_optimized]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][memory_mb]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][memory_mb]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][num_cores]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][num_cores]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][disk_gb]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][disk_gb]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][automatic_instance_store_mapping]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][automatic_instance_store_mapping]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_performance]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_performance]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_type_uid]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_type_uid]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][iam_instance_profile]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][iam_instance_profile]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_size]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][root_volume_size]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][num_cores]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][num_cores]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][memory_mb]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][memory_mb]
						// Discovered as a child of instance at key cloud_specific_attributes
						// My children became the parent map after being discovered as a child of instance at key cloud_specific_attributes -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][disk_gb]
						// (At Root) My children became the parent map. -- Submatch: ["instance[cloud_specific_attributes]" "instance" "cloud_specific_attributes"] -- CmdFlag: instance[cloud_specific_attributes][disk_gb]
						"cloud_specific_attributes": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// DEBUG INFO Created for key automatic_instance_store_mapping. -- Submatch: ["[automatic_instance_store_mapping]" "" "automatic_instance_store_mapping"] at idx 1 -- cmdFlagName: instance[cloud_specific_attributes][automatic_instance_store_mapping] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at automatic_instance_store_mapping because temp_attr wasn't null, which it never will be
									// Discovered at key automatic_instance_store_mapping of parent map and updating -- Submatch: ["[automatic_instance_store_mapping]" "" "automatic_instance_store_mapping"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at automatic_instance_store_mapping because temp_attr wasn't null, which it never will be
									"automatic_instance_store_mapping": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key disk_gb. -- Submatch: ["[disk_gb]" "" "disk_gb"] at idx 1 -- cmdFlagName: instance[cloud_specific_attributes][disk_gb] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at disk_gb because temp_attr wasn't null, which it never will be
									// Discovered at key disk_gb of parent map and updating -- Submatch: ["[disk_gb]" "" "disk_gb"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at disk_gb because temp_attr wasn't null, which it never will be
									"disk_gb": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									// DEBUG INFO Created for key ebs_optimized. -- Submatch: ["[ebs_optimized]" "" "ebs_optimized"] at idx 1 -- cmdFlagName: instance[cloud_specific_attributes][ebs_optimized] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at ebs_optimized because temp_attr wasn't null, which it never will be
									"ebs_optimized": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key iam_instance_profile. -- Submatch: ["[iam_instance_profile]" "" "iam_instance_profile"] at idx 1 -- cmdFlagName: instance[cloud_specific_attributes][iam_instance_profile] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at iam_instance_profile because temp_attr wasn't null, which it never will be
									// Discovered at key iam_instance_profile of parent map and updating -- Submatch: ["[iam_instance_profile]" "" "iam_instance_profile"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at iam_instance_profile because temp_attr wasn't null, which it never will be
									"iam_instance_profile": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key memory_mb. -- Submatch: ["[memory_mb]" "" "memory_mb"] at idx 1 -- cmdFlagName: instance[cloud_specific_attributes][memory_mb] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at memory_mb because temp_attr wasn't null, which it never will be
									// Discovered at key memory_mb of parent map and updating -- Submatch: ["[memory_mb]" "" "memory_mb"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at memory_mb because temp_attr wasn't null, which it never will be
									"memory_mb": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									// DEBUG INFO Created for key num_cores. -- Submatch: ["[num_cores]" "" "num_cores"] at idx 1 -- cmdFlagName: instance[cloud_specific_attributes][num_cores] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at num_cores because temp_attr wasn't null, which it never will be
									// Discovered at key num_cores of parent map and updating -- Submatch: ["[num_cores]" "" "num_cores"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at num_cores because temp_attr wasn't null, which it never will be
									"num_cores": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									// DEBUG INFO Created for key root_volume_performance. -- Submatch: ["[root_volume_performance]" "" "root_volume_performance"] at idx 1 -- cmdFlagName: instance[cloud_specific_attributes][root_volume_performance] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at root_volume_performance because temp_attr wasn't null, which it never will be
									// Discovered at key root_volume_performance of parent map and updating -- Submatch: ["[root_volume_performance]" "" "root_volume_performance"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at root_volume_performance because temp_attr wasn't null, which it never will be
									"root_volume_performance": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key root_volume_size. -- Submatch: ["[root_volume_size]" "" "root_volume_size"] at idx 1 -- cmdFlagName: instance[cloud_specific_attributes][root_volume_size] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at root_volume_size because temp_attr wasn't null, which it never will be
									// Discovered at key root_volume_size of parent map and updating -- Submatch: ["[root_volume_size]" "" "root_volume_size"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at root_volume_size because temp_attr wasn't null, which it never will be
									"root_volume_size": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key root_volume_type_uid. -- Submatch: ["[root_volume_type_uid]" "" "root_volume_type_uid"] at idx 1 -- cmdFlagName: instance[cloud_specific_attributes][root_volume_type_uid] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at root_volume_type_uid because temp_attr wasn't null, which it never will be
									// Discovered at key root_volume_type_uid of parent map and updating -- Submatch: ["[root_volume_type_uid]" "" "root_volume_type_uid"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at root_volume_type_uid because temp_attr wasn't null, which it never will be
									"root_volume_type_uid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						// DEBUG INFO Child at parent[instance][security_group_hrefs] -- Submatch: ["instance[security_group_hrefs]" "instance" "security_group_hrefs"] at idx 0 -- cmdFlagName: instance[security_group_hrefs][]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["instance[security_group_hrefs]" "instance" "security_group_hrefs"] -- CmdFlag: instance[security_group_hrefs][]
						// Discovered as a child of instance at key security_group_hrefs
						// My children became the parent map after being discovered as a child of instance at key security_group_hrefs -- Submatch: ["instance[security_group_hrefs]" "instance" "security_group_hrefs"] -- CmdFlag: instance[security_group_hrefs][]
						// (At Root) My children became the parent map. -- Submatch: ["instance[security_group_hrefs]" "instance" "security_group_hrefs"] -- CmdFlag: instance[security_group_hrefs][]
						"security_group_hrefs": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						// DEBUG INFO Child at parent[instance][subnet_hrefs] -- Submatch: ["instance[subnet_hrefs]" "instance" "subnet_hrefs"] at idx 0 -- cmdFlagName: instance[subnet_hrefs][]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["instance[subnet_hrefs]" "instance" "subnet_hrefs"] -- CmdFlag: instance[subnet_hrefs][]
						// Discovered as a child of instance at key subnet_hrefs
						// My children became the parent map after being discovered as a child of instance at key subnet_hrefs -- Submatch: ["instance[subnet_hrefs]" "instance" "subnet_hrefs"] -- CmdFlag: instance[subnet_hrefs][]
						// (At Root) My children became the parent map. -- Submatch: ["instance[subnet_hrefs]" "instance" "subnet_hrefs"] -- CmdFlag: instance[subnet_hrefs][]
						"subnet_hrefs": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			// DEBUG INFO Created for key api_behavior. -- Submatch: ["api_behavior" "api_behavior" ""] at idx 0 -- cmdFlagName: api_behavior -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at api_behavior because temp_attr wasn't null, which it never will be
			"api_behavior": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			// DEBUG INFO
			"locked": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"user_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"inherited_sources": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"os_platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"security_groups": &schema.Schema{
				Type:     schema.TypeList, //[]SecurityGroup,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"public_ip_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			// DEBUG INFO
			"inputs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"terminated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"monitoring_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"monitoring_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"pricing_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"public_dns_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			// DEBUG INFO
			"private_dns_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			// DEBUG INFO
			"subnets": &schema.Schema{
				Type:     schema.TypeList, //[]Subnet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"associate_public_ip_address": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"ip_forwarding_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"admin_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Instance", "create", "/api/clouds/%s/instances", params)
	if err != nil {
		message := fmt.Sprintf("Instance Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Instance Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Instance HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Instance Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleInstanceRead(d, meta)
	return nil
}

func resourceRightScaleInstanceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleInstanceCustomLodgement() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names quantity,timeframe
		// Pattern:/api/clouds/%s/instances/%s/instance_custom_lodgements Vars:cloud_id,instance_id

		Create: resourceRightScaleInstanceCustomLodgementCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleInstanceCustomLodgementDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleInstanceCustomLodgementRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names quantity
		// Pattern:/api/clouds/%s/instances/%s/instance_custom_lodgements/%s Vars:cloud_id,instance_id,id

		Update: resourceRightScaleInstanceCustomLodgementUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key timeframe. -- Submatch: ["timeframe" "timeframe" ""] at idx 0 -- cmdFlagName: timeframe -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at timeframe because temp_attr wasn't null, which it never will be
			"timeframe": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"end_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"resource_instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_launched_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"start_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"account_owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_billing_start_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_template_library_href": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_template_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_template_published_by_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key quantity. -- Submatch: ["quantity[]" "quantity" ""] at idx 0 -- cmdFlagName: quantity[][value] -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at quantity because temp_attr wasn't null, which it never will be
			// Discovered at key quantity of parent map and updating -- Submatch: ["quantity[]" "quantity" ""] at idx 0 -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at quantity because temp_attr wasn't null, which it never will be
			// Discovered at key quantity of parent map and updating -- Submatch: ["quantity[]" "quantity" ""] at idx 0 -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at quantity because temp_attr wasn't null, which it never will be
			// Discovered at key quantity of parent map and updating -- Submatch: ["quantity[]" "quantity" ""] at idx 0 -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at quantity because temp_attr wasn't null, which it never will be
			"quantity": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key value. -- Submatch: ["[value]" "" "value"] at idx 1 -- cmdFlagName: quantity[][value] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at value because temp_attr wasn't null, which it never will be
			// Discovered at key value of parent map and updating -- Submatch: ["[value]" "" "value"] at idx 1 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at value because temp_attr wasn't null, which it never will be
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			// DEBUG INFO Created for key name. -- Submatch: ["[name]" "" "name"] at idx 1 -- cmdFlagName: quantity[][name] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at name because temp_attr wasn't null, which it never will be
			// Discovered at key name of parent map and updating -- Submatch: ["[name]" "" "name"] at idx 1 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at name because temp_attr wasn't null, which it never will be
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"resource_billing_end_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleInstanceCustomLodgementCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("InstanceCustomLodgement", "create", "/api/clouds/%s/instances/%s/instance_custom_lodgements", params)
	if err != nil {
		message := fmt.Sprintf("InstanceCustomLodgement Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("InstanceCustomLodgement Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("InstanceCustomLodgement HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("InstanceCustomLodgement Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleInstanceCustomLodgementRead(d, meta)
	return nil
}
func resourceRightScaleInstanceCustomLodgementDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleInstanceCustomLodgementRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleInstanceCustomLodgementUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleInstanceType() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleInstanceTypeRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"cpu_speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"local_disk_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"memory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"cpu_architecture": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"cpu_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"local_disks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleInstanceTypeRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleIpAddress() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names ip_address
		// Pattern:/api/clouds/%s/ip_addresses Vars:cloud_id

		Create: resourceRightScaleIpAddressCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleIpAddressDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleIpAddressRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names ip_address
		// Pattern:/api/clouds/%s/ip_addresses/%s Vars:cloud_id,id

		Update: resourceRightScaleIpAddressUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key ip_address. -- Submatch: ["ip_address[deployment_href]" "ip_address" "deployment_href"] at idx 0 -- cmdFlagName: ip_address[deployment_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address because temp_attr wasn't null, which it never will be
			// Discovered at key ip_address of parent map and updating -- Submatch: ["ip_address[network_href]" "ip_address" "network_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address because temp_attr wasn't null, which it never will be
			// Discovered at key ip_address of parent map and updating -- Submatch: ["ip_address[domain]" "ip_address" "domain"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address because temp_attr wasn't null, which it never will be
			// Discovered at key ip_address of parent map and updating -- Submatch: ["ip_address[name]" "ip_address" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address because temp_attr wasn't null, which it never will be
			// Discovered at key ip_address of parent map and updating -- Submatch: ["ip_address[deployment_href]" "ip_address" "deployment_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address because temp_attr wasn't null, which it never will be
			// Discovered at key ip_address of parent map and updating -- Submatch: ["ip_address[name]" "ip_address" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address because temp_attr wasn't null, which it never will be
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleIpAddressCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("IpAddress", "create", "/api/clouds/%s/ip_addresses", params)
	if err != nil {
		message := fmt.Sprintf("IpAddress Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("IpAddress Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("IpAddress HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("IpAddress Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleIpAddressRead(d, meta)
	return nil
}
func resourceRightScaleIpAddressDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleIpAddressRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleIpAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleIpAddressBinding() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names ip_address_binding
		// Pattern:/api/clouds/%s/ip_addresses/%s/ip_address_bindings Vars:cloud_id,ip_address_id
		// Pattern:/api/clouds/%s/ip_address_bindings Vars:cloud_id

		Create: resourceRightScaleIpAddressBindingCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleIpAddressBindingDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleIpAddressBindingRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key ip_address_binding. -- Submatch: ["ip_address_binding[public_ip_address_href]" "ip_address_binding" "public_ip_address_href"] at idx 0 -- cmdFlagName: ip_address_binding[public_ip_address_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address_binding because temp_attr wasn't null, which it never will be
			// Discovered at key ip_address_binding of parent map and updating -- Submatch: ["ip_address_binding[instance_href]" "ip_address_binding" "instance_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address_binding because temp_attr wasn't null, which it never will be
			// Discovered at key ip_address_binding of parent map and updating -- Submatch: ["ip_address_binding[private_port]" "ip_address_binding" "private_port"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address_binding because temp_attr wasn't null, which it never will be
			// Discovered at key ip_address_binding of parent map and updating -- Submatch: ["ip_address_binding[public_port]" "ip_address_binding" "public_port"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address_binding because temp_attr wasn't null, which it never will be
			// Discovered at key ip_address_binding of parent map and updating -- Submatch: ["ip_address_binding[protocol]" "ip_address_binding" "protocol"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ip_address_binding because temp_attr wasn't null, which it never will be
			"ip_address_binding": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"private_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"public_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"recurring": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleIpAddressBindingCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("IpAddressBinding", "create", "/api/clouds/%s/ip_addresses/%s/ip_address_bindings", params)
	if err != nil {
		message := fmt.Sprintf("IpAddressBinding Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("IpAddressBinding Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("IpAddressBinding HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("IpAddressBinding Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleIpAddressBindingRead(d, meta)
	return nil
}
func resourceRightScaleIpAddressBindingDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleIpAddressBindingRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleMonitoringMetric() *schema.Resource {
	return &schema.Resource{
		// ACTION: data
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names end,start

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names period,size,title,tz

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names period,size,title,tz

		Read: resourceRightScaleMonitoringMetricRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"graph_href": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"plugin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"view": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleMonitoringMetricRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleMultiCloudImage() *schema.Resource {
	return &schema.Resource{
		// ACTION: clone
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names multi_cloud_image

		// ACTION: commit
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names commit_message

		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names multi_cloud_image
		// Pattern:/api/server_templates/%s/multi_cloud_images Vars:server_template_id
		// Pattern:/api/multi_cloud_images Vars:

		Create: resourceRightScaleMultiCloudImageCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleMultiCloudImageDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleMultiCloudImageRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names multi_cloud_image
		// Pattern:/api/server_templates/%s/multi_cloud_images/%s Vars:server_template_id,id
		// Pattern:/api/multi_cloud_images/%s Vars:id

		Update: resourceRightScaleMultiCloudImageUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key multi_cloud_image. -- Submatch: ["multi_cloud_image[description]" "multi_cloud_image" "description"] at idx 0 -- cmdFlagName: multi_cloud_image[description] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image of parent map and updating -- Submatch: ["multi_cloud_image[name]" "multi_cloud_image" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image of parent map and updating -- Submatch: ["multi_cloud_image[description]" "multi_cloud_image" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image of parent map and updating -- Submatch: ["multi_cloud_image[name]" "multi_cloud_image" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image because temp_attr wasn't null, which it never will be
			"multi_cloud_image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleMultiCloudImageCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("MultiCloudImage", "create", "/api/server_templates/%s/multi_cloud_images", params)
	if err != nil {
		message := fmt.Sprintf("MultiCloudImage Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("MultiCloudImage Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("MultiCloudImage HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("MultiCloudImage Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleMultiCloudImageRead(d, meta)
	return nil
}
func resourceRightScaleMultiCloudImageDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleMultiCloudImageRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleMultiCloudImageUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleMultiCloudImageSetting() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names multi_cloud_image_setting
		// Pattern:/api/multi_cloud_images/%s/settings Vars:multi_cloud_image_id

		Create: resourceRightScaleMultiCloudImageSettingCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleMultiCloudImageSettingDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleMultiCloudImageSettingRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names multi_cloud_image_setting
		// Pattern:/api/multi_cloud_images/%s/settings/%s Vars:multi_cloud_image_id,id

		Update: resourceRightScaleMultiCloudImageSettingUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key multi_cloud_image_setting. -- Submatch: ["multi_cloud_image_setting[instance_type_href]" "multi_cloud_image_setting" "instance_type_href"] at idx 0 -- cmdFlagName: multi_cloud_image_setting[instance_type_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[ramdisk_image_href]" "multi_cloud_image_setting" "ramdisk_image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[kernel_image_href]" "multi_cloud_image_setting" "kernel_image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[cloud_href]" "multi_cloud_image_setting" "cloud_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[image_href]" "multi_cloud_image_setting" "image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[user_data]" "multi_cloud_image_setting" "user_data"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[instance_type_href]" "multi_cloud_image_setting" "instance_type_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[ramdisk_image_href]" "multi_cloud_image_setting" "ramdisk_image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[kernel_image_href]" "multi_cloud_image_setting" "kernel_image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[cloud_href]" "multi_cloud_image_setting" "cloud_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[image_href]" "multi_cloud_image_setting" "image_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			// Discovered at key multi_cloud_image_setting of parent map and updating -- Submatch: ["multi_cloud_image_setting[user_data]" "multi_cloud_image_setting" "user_data"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at multi_cloud_image_setting because temp_attr wasn't null, which it never will be
			"multi_cloud_image_setting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleMultiCloudImageSettingCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("MultiCloudImageSetting", "create", "/api/multi_cloud_images/%s/settings", params)
	if err != nil {
		message := fmt.Sprintf("MultiCloudImageSetting Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("MultiCloudImageSetting Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("MultiCloudImageSetting HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("MultiCloudImageSetting Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleMultiCloudImageSettingRead(d, meta)
	return nil
}
func resourceRightScaleMultiCloudImageSettingDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleMultiCloudImageSettingRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleMultiCloudImageSettingUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleNetwork() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names network
		// Pattern:/api/networks Vars:

		Create: resourceRightScaleNetworkCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleNetworkDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleNetworkRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names network
		// Pattern:/api/networks/%s Vars:id

		Update: resourceRightScaleNetworkUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"cidr_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"instance_tenancy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key network. -- Submatch: ["network[instance_tenancy]" "network" "instance_tenancy"] at idx 0 -- cmdFlagName: network[instance_tenancy] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network because temp_attr wasn't null, which it never will be
			// Discovered at key network of parent map and updating -- Submatch: ["network[description]" "network" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network because temp_attr wasn't null, which it never will be
			// Discovered at key network of parent map and updating -- Submatch: ["network[cidr_block]" "network" "cidr_block"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network because temp_attr wasn't null, which it never will be
			// Discovered at key network of parent map and updating -- Submatch: ["network[cloud_href]" "network" "cloud_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network because temp_attr wasn't null, which it never will be
			// Discovered at key network of parent map and updating -- Submatch: ["network[name]" "network" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network because temp_attr wasn't null, which it never will be
			// Discovered at key network of parent map and updating -- Submatch: ["network[route_table_href]" "network" "route_table_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network because temp_attr wasn't null, which it never will be
			// Discovered at key network of parent map and updating -- Submatch: ["network[description]" "network" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network because temp_attr wasn't null, which it never will be
			// Discovered at key network of parent map and updating -- Submatch: ["network[name]" "network" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network because temp_attr wasn't null, which it never will be
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Network", "create", "/api/networks", params)
	if err != nil {
		message := fmt.Sprintf("Network Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Network Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Network HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Network Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleNetworkRead(d, meta)
	return nil
}
func resourceRightScaleNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleNetworkRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleNetworkGateway() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names network_gateway
		// Pattern:/api/network_gateways Vars:

		Create: resourceRightScaleNetworkGatewayCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleNetworkGatewayDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleNetworkGatewayRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names network_gateway
		// Pattern:/api/network_gateways/%s Vars:id

		Update: resourceRightScaleNetworkGatewayUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key network_gateway. -- Submatch: ["network_gateway[description]" "network_gateway" "description"] at idx 0 -- cmdFlagName: network_gateway[description] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_gateway because temp_attr wasn't null, which it never will be
			// Discovered at key network_gateway of parent map and updating -- Submatch: ["network_gateway[cloud_href]" "network_gateway" "cloud_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_gateway because temp_attr wasn't null, which it never will be
			// Discovered at key network_gateway of parent map and updating -- Submatch: ["network_gateway[name]" "network_gateway" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_gateway because temp_attr wasn't null, which it never will be
			// Discovered at key network_gateway of parent map and updating -- Submatch: ["network_gateway[type]" "network_gateway" "type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_gateway because temp_attr wasn't null, which it never will be
			// Discovered at key network_gateway of parent map and updating -- Submatch: ["network_gateway[network_href]" "network_gateway" "network_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_gateway because temp_attr wasn't null, which it never will be
			// Discovered at key network_gateway of parent map and updating -- Submatch: ["network_gateway[description]" "network_gateway" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_gateway because temp_attr wasn't null, which it never will be
			// Discovered at key network_gateway of parent map and updating -- Submatch: ["network_gateway[name]" "network_gateway" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_gateway because temp_attr wasn't null, which it never will be
			"network_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleNetworkGatewayCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("NetworkGateway", "create", "/api/network_gateways", params)
	if err != nil {
		message := fmt.Sprintf("NetworkGateway Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("NetworkGateway Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("NetworkGateway HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("NetworkGateway Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleNetworkGatewayRead(d, meta)
	return nil
}
func resourceRightScaleNetworkGatewayDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleNetworkGatewayRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleNetworkGatewayUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleNetworkOptionGroup() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names network_option_group
		// Pattern:/api/network_option_groups Vars:

		Create: resourceRightScaleNetworkOptionGroupCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleNetworkOptionGroupDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleNetworkOptionGroupRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names network_option_group
		// Pattern:/api/network_option_groups/%s Vars:id

		Update: resourceRightScaleNetworkOptionGroupUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key network_option_group. -- Submatch: ["network_option_group[description]" "network_option_group" "description"] at idx 0 -- cmdFlagName: network_option_group[description] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group because temp_attr wasn't null, which it never will be
			// Discovered at key network_option_group of parent map and updating -- Submatch: ["network_option_group[cloud_href]" "network_option_group" "cloud_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group because temp_attr wasn't null, which it never will be
			// Discovered at key network_option_group of parent map and updating -- Submatch: ["network_option_group[options]" "network_option_group" "options"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group because temp_attr wasn't null, which it never will be
			// Discovered at key network_option_group of parent map and updating -- Submatch: ["network_option_group[type]" "network_option_group" "type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group because temp_attr wasn't null, which it never will be
			// Discovered at key network_option_group of parent map and updating -- Submatch: ["network_option_group[name]" "network_option_group" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group because temp_attr wasn't null, which it never will be
			// Discovered at key network_option_group of parent map and updating -- Submatch: ["network_option_group[description]" "network_option_group" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group because temp_attr wasn't null, which it never will be
			// Discovered at key network_option_group of parent map and updating -- Submatch: ["network_option_group[options]" "network_option_group" "options"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group because temp_attr wasn't null, which it never will be
			// Discovered at key network_option_group of parent map and updating -- Submatch: ["network_option_group[name]" "network_option_group" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group because temp_attr wasn't null, which it never will be
			"network_option_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"options": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleNetworkOptionGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("NetworkOptionGroup", "create", "/api/network_option_groups", params)
	if err != nil {
		message := fmt.Sprintf("NetworkOptionGroup Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("NetworkOptionGroup Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("NetworkOptionGroup HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("NetworkOptionGroup Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleNetworkOptionGroupRead(d, meta)
	return nil
}
func resourceRightScaleNetworkOptionGroupDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleNetworkOptionGroupRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleNetworkOptionGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleNetworkOptionGroupAttachment() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names network_option_group_attachment
		// Pattern:/api/network_option_group_attachments Vars:

		Create: resourceRightScaleNetworkOptionGroupAttachmentCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleNetworkOptionGroupAttachmentDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleNetworkOptionGroupAttachmentRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names network_option_group_attachment
		// Pattern:/api/network_option_group_attachments/%s Vars:id

		Update: resourceRightScaleNetworkOptionGroupAttachmentUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"network_option_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key network_option_group_attachment. -- Submatch: ["network_option_group_attachment[network_option_group_href]" "network_option_group_attachment" "network_option_group_href"] at idx 0 -- cmdFlagName: network_option_group_attachment[network_option_group_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key network_option_group_attachment of parent map and updating -- Submatch: ["network_option_group_attachment[network_href]" "network_option_group_attachment" "network_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key network_option_group_attachment of parent map and updating -- Submatch: ["network_option_group_attachment[network_option_group_href]" "network_option_group_attachment" "network_option_group_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at network_option_group_attachment because temp_attr wasn't null, which it never will be
			"network_option_group_attachment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleNetworkOptionGroupAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("NetworkOptionGroupAttachment", "create", "/api/network_option_group_attachments", params)
	if err != nil {
		message := fmt.Sprintf("NetworkOptionGroupAttachment Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("NetworkOptionGroupAttachment Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("NetworkOptionGroupAttachment HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("NetworkOptionGroupAttachment Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleNetworkOptionGroupAttachmentRead(d, meta)
	return nil
}
func resourceRightScaleNetworkOptionGroupAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleNetworkOptionGroupAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleNetworkOptionGroupAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleOauth2() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names account_id,client_id,client_secret,grant_type,r_s_version,refresh_token,right_link_version
		// Pattern:/api/oauth2/ Vars:

		Create: resourceRightScaleOauth2Create,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key r_s_version. -- Submatch: ["r_s_version" "r_s_version" ""] at idx 0 -- cmdFlagName: r_s_version -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at r_s_version because temp_attr wasn't null, which it never will be
			"r_s_version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO Created for key grant_type. -- Submatch: ["grant_type" "grant_type" ""] at idx 0 -- cmdFlagName: grant_type -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at grant_type because temp_attr wasn't null, which it never will be
			"grant_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// DEBUG INFO Created for key account_id. -- Submatch: ["account_id" "account_id" ""] at idx 0 -- cmdFlagName: account_id -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at account_id because temp_attr wasn't null, which it never will be
			"account_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO Created for key client_id. -- Submatch: ["client_id" "client_id" ""] at idx 0 -- cmdFlagName: client_id -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at client_id because temp_attr wasn't null, which it never will be
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO Created for key right_link_version. -- Submatch: ["right_link_version" "right_link_version" ""] at idx 0 -- cmdFlagName: right_link_version -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at right_link_version because temp_attr wasn't null, which it never will be
			"right_link_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO Created for key client_secret. -- Submatch: ["client_secret" "client_secret" ""] at idx 0 -- cmdFlagName: client_secret -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at client_secret because temp_attr wasn't null, which it never will be
			"client_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO Created for key refresh_token. -- Submatch: ["refresh_token" "refresh_token" ""] at idx 0 -- cmdFlagName: refresh_token -- Operating at root: true
			// Matched at subMatch[2] == '' -- Operating at root: true
			// Assigned to parent map at refresh_token because temp_attr wasn't null, which it never will be
			"refresh_token": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRightScaleOauth2Create(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Oauth2", "create", "/api/oauth2/", params)
	if err != nil {
		message := fmt.Sprintf("Oauth2 Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Oauth2 Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Oauth2 HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Oauth2 Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleOauth2Read(d, meta)
	return nil
}

func resourceRightScalePermission() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names permission
		// Pattern:/api/permissions Vars:

		Create: resourceRightScalePermissionCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScalePermissionDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScalePermissionRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key permission. -- Submatch: ["permission[role_title]" "permission" "role_title"] at idx 0 -- cmdFlagName: permission[role_title] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at permission because temp_attr wasn't null, which it never will be
			// Discovered at key permission of parent map and updating -- Submatch: ["permission[user_href]" "permission" "user_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at permission because temp_attr wasn't null, which it never will be
			"permission": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"role_title": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScalePermissionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Permission", "create", "/api/permissions", params)
	if err != nil {
		message := fmt.Sprintf("Permission Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Permission Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Permission HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Permission Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScalePermissionRead(d, meta)
	return nil
}
func resourceRightScalePermissionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScalePermissionRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScalePlacementGroup() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names placement_group
		// Pattern:/api/placement_groups Vars:

		Create: resourceRightScalePlacementGroupCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScalePlacementGroupDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScalePlacementGroupRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key placement_group. -- Submatch: ["placement_group[description]" "placement_group" "description"] at idx 0 -- cmdFlagName: placement_group[description] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at placement_group because temp_attr wasn't null, which it never will be
			// Discovered at key placement_group of parent map and updating -- Submatch: ["placement_group[cloud_href]" "placement_group" "cloud_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at placement_group because temp_attr wasn't null, which it never will be
			// Discovered at key placement_group of parent map and updating -- Submatch: ["placement_group[name]" "placement_group" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at placement_group because temp_attr wasn't null, which it never will be
			"placement_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScalePlacementGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("PlacementGroup", "create", "/api/placement_groups", params)
	if err != nil {
		message := fmt.Sprintf("PlacementGroup Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("PlacementGroup Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("PlacementGroup HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("PlacementGroup Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScalePlacementGroupRead(d, meta)
	return nil
}
func resourceRightScalePlacementGroupDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScalePlacementGroupRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScalePreference() *schema.Resource {
	return &schema.Resource{
		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScalePreferenceDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScalePreferenceRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names preference
		// Pattern:/api/preferences/%s Vars:id

		Update: resourceRightScalePreferenceUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key preference. -- Submatch: ["preference[contents]" "preference" "contents"] at idx 0 -- cmdFlagName: preference[contents] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at preference because temp_attr wasn't null, which it never will be
			"preference": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"contents": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScalePreferenceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScalePreferenceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScalePreferenceUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScalePublication() *schema.Resource {
	return &schema.Resource{
		// ACTION: import
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScalePublicationRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"publisher": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"revision_notes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"commit_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"content_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScalePublicationRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScalePublicationLineage() *schema.Resource {
	return &schema.Resource{
		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScalePublicationLineageRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"short_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"comments_emailed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"comments_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"long_description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScalePublicationLineageRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRecurringVolumeAttachment() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names recurring_volume_attachment
		// Pattern:/api/clouds/%s/recurring_volume_attachments Vars:cloud_id
		// Pattern:/api/clouds/%s/volumes/%s/recurring_volume_attachments Vars:cloud_id,volume_id
		// Pattern:/api/clouds/%s/volume_snapshots/%s/recurring_volume_attachments Vars:cloud_id,volume_snapshot_id

		Create: resourceRightScaleRecurringVolumeAttachmentCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleRecurringVolumeAttachmentDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleRecurringVolumeAttachmentRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"runnable_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"device_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"storage_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key recurring_volume_attachment. -- Submatch: ["recurring_volume_attachment[volume_type_href]" "recurring_volume_attachment" "volume_type_href"] at idx 0 -- cmdFlagName: recurring_volume_attachment[volume_type_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at recurring_volume_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key recurring_volume_attachment of parent map and updating -- Submatch: ["recurring_volume_attachment[runnable_href]" "recurring_volume_attachment" "runnable_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at recurring_volume_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key recurring_volume_attachment of parent map and updating -- Submatch: ["recurring_volume_attachment[storage_href]" "recurring_volume_attachment" "storage_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at recurring_volume_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key recurring_volume_attachment of parent map and updating -- Submatch: ["recurring_volume_attachment[settings]" "recurring_volume_attachment" "settings"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at recurring_volume_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key recurring_volume_attachment of parent map and updating -- Submatch: ["recurring_volume_attachment[device]" "recurring_volume_attachment" "device"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at recurring_volume_attachment because temp_attr wasn't null, which it never will be
			"recurring_volume_attachment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRightScaleRecurringVolumeAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("RecurringVolumeAttachment", "create", "/api/clouds/%s/recurring_volume_attachments", params)
	if err != nil {
		message := fmt.Sprintf("RecurringVolumeAttachment Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("RecurringVolumeAttachment Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("RecurringVolumeAttachment HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("RecurringVolumeAttachment Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleRecurringVolumeAttachmentRead(d, meta)
	return nil
}
func resourceRightScaleRecurringVolumeAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRecurringVolumeAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRepository() *schema.Resource {
	return &schema.Resource{
		// ACTION: cookbook_import
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names asset_hrefs,follow,namespace,repository_commit_reference,with_dependencies

		// ACTION: cookbook_import_preview
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names asset_hrefs,namespace

		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names repository
		// Pattern:/api/repositories Vars:

		Create: resourceRightScaleRepositoryCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleRepositoryDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: refetch
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names auto_import

		// ACTION: resolve
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names imported_cookbook_name

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleRepositoryRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names repository
		// Pattern:/api/repositories/%s Vars:id

		Update: resourceRightScaleRepositoryUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"asset_counts": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"commit_reference": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"fetch_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created as a child node when an existing child of the same key was not found. Parent was repository -- Submatch: ["repository[credentials]" "repository" "credentials"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[credentials]" "repository" "credentials"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[credentials]" "repository" "credentials"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[commit_reference]" "repository" "commit_reference"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[source_type]" "repository" "source_type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[auto_import]" "repository" "auto_import"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[description]" "repository" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[source]" "repository" "source"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[name]" "repository" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[asset_paths]" "repository" "asset_paths"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[credentials]" "repository" "credentials"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[credentials]" "repository" "credentials"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[credentials]" "repository" "credentials"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[commit_reference]" "repository" "commit_reference"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[description]" "repository" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[source_type]" "repository" "source_type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[source]" "repository" "source"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			// Discovered at key repository of parent map and updating -- Submatch: ["repository[name]" "repository" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at repository because temp_attr wasn't null, which it never will be
			"repository": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// DEBUG INFO Child at parent[repository][asset_paths] -- Submatch: ["repository[asset_paths]" "repository" "asset_paths"] at idx 0 -- cmdFlagName: repository[asset_paths][cookbooks][]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["repository[asset_paths]" "repository" "asset_paths"] -- CmdFlag: repository[asset_paths][cookbooks][]
						// Discovered as a child of repository at key asset_paths
						// My children became the parent map after being discovered as a child of repository at key asset_paths -- Submatch: ["repository[asset_paths]" "repository" "asset_paths"] -- CmdFlag: repository[asset_paths][cookbooks][]
						// (At Root) My children became the parent map. -- Submatch: ["repository[asset_paths]" "repository" "asset_paths"] -- CmdFlag: repository[asset_paths][cookbooks][]
						"asset_paths": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						// DEBUG INFO Child at parent[repository][credentials] -- Submatch: ["repository[credentials]" "repository" "credentials"] at idx 0 -- cmdFlagName: repository[credentials][password]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][password]
						// Discovered as a child of repository at key credentials
						// My children became the parent map after being discovered as a child of repository at key credentials -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][username]
						// (At Root) My children became the parent map. -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][username]
						// Discovered as a child of repository at key credentials
						// My children became the parent map after being discovered as a child of repository at key credentials -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][ssh_key]
						// (At Root) My children became the parent map. -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][ssh_key]
						// Discovered as a child of repository at key credentials
						// My children became the parent map after being discovered as a child of repository at key credentials -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][username]
						// (At Root) My children became the parent map. -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][username]
						// Discovered as a child of repository at key credentials
						// My children became the parent map after being discovered as a child of repository at key credentials -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][password]
						// (At Root) My children became the parent map. -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][password]
						// Discovered as a child of repository at key credentials
						// My children became the parent map after being discovered as a child of repository at key credentials -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][ssh_key]
						// (At Root) My children became the parent map. -- Submatch: ["repository[credentials]" "repository" "credentials"] -- CmdFlag: repository[credentials][ssh_key]
						"credentials": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// DEBUG INFO Created for key password. -- Submatch: ["[password]" "" "password"] at idx 1 -- cmdFlagName: repository[credentials][password] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at password because temp_attr wasn't null, which it never will be
									// Discovered at key password of parent map and updating -- Submatch: ["[password]" "" "password"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at password because temp_attr wasn't null, which it never will be
									"password": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key ssh_key. -- Submatch: ["[ssh_key]" "" "ssh_key"] at idx 1 -- cmdFlagName: repository[credentials][ssh_key] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at ssh_key because temp_attr wasn't null, which it never will be
									// Discovered at key ssh_key of parent map and updating -- Submatch: ["[ssh_key]" "" "ssh_key"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at ssh_key because temp_attr wasn't null, which it never will be
									"ssh_key": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key username. -- Submatch: ["[username]" "" "username"] at idx 1 -- cmdFlagName: repository[credentials][username] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at username because temp_attr wasn't null, which it never will be
									// Discovered at key username of parent map and updating -- Submatch: ["[username]" "" "username"] at idx 1 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at username because temp_attr wasn't null, which it never will be
									"username": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"source_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Repository", "create", "/api/repositories", params)
	if err != nil {
		message := fmt.Sprintf("Repository Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Repository Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Repository HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Repository Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleRepositoryRead(d, meta)
	return nil
}
func resourceRightScaleRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRepositoryAsset() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleRepositoryAssetRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleRepositoryAssetRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRightScript() *schema.Resource {
	return &schema.Resource{
		// ACTION: commit
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names right_script

		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names right_script
		// Pattern:/api/right_scripts Vars:

		Create: resourceRightScaleRightScriptCreate,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names latest_only

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleRightScriptRead,

		// ACTION: show_source
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names right_script
		// Pattern:/api/right_scripts/%s Vars:id

		Update: resourceRightScaleRightScriptUpdate,

		// ACTION: update_source
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names filename

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"lineage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key right_script. -- Submatch: ["right_script[description]" "right_script" "description"] at idx 0 -- cmdFlagName: right_script[description] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script because temp_attr wasn't null, which it never will be
			// Discovered at key right_script of parent map and updating -- Submatch: ["right_script[source]" "right_script" "source"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script because temp_attr wasn't null, which it never will be
			// Discovered at key right_script of parent map and updating -- Submatch: ["right_script[name]" "right_script" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script because temp_attr wasn't null, which it never will be
			// Discovered at key right_script of parent map and updating -- Submatch: ["right_script[description]" "right_script" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script because temp_attr wasn't null, which it never will be
			// Discovered at key right_script of parent map and updating -- Submatch: ["right_script[source]" "right_script" "source"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script because temp_attr wasn't null, which it never will be
			// Discovered at key right_script of parent map and updating -- Submatch: ["right_script[name]" "right_script" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script because temp_attr wasn't null, which it never will be
			"right_script": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"inputs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleRightScriptCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("RightScript", "create", "/api/right_scripts", params)
	if err != nil {
		message := fmt.Sprintf("RightScript Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("RightScript Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("RightScript HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("RightScript Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleRightScriptRead(d, meta)
	return nil
}

func resourceRightScaleRightScriptRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRightScriptUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRightScriptAttachment() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names right_script_attachment
		// Pattern:/api/right_scripts/%s/attachments Vars:right_script_id

		Create: resourceRightScaleRightScriptAttachmentCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleRightScriptAttachmentDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleRightScriptAttachmentRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names right_script_attachment
		// Pattern:/api/right_scripts/%s/attachments/%s Vars:right_script_id,id

		Update: resourceRightScaleRightScriptAttachmentUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key right_script_attachment. -- Submatch: ["right_script_attachment[filename]" "right_script_attachment" "filename"] at idx 0 -- cmdFlagName: right_script_attachment[filename] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key right_script_attachment of parent map and updating -- Submatch: ["right_script_attachment[content]" "right_script_attachment" "content"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key right_script_attachment of parent map and updating -- Submatch: ["right_script_attachment[filename]" "right_script_attachment" "filename"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key right_script_attachment of parent map and updating -- Submatch: ["right_script_attachment[content]" "right_script_attachment" "content"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at right_script_attachment because temp_attr wasn't null, which it never will be
			"right_script_attachment": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"digest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleRightScriptAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("RightScriptAttachment", "create", "/api/right_scripts/%s/attachments", params)
	if err != nil {
		message := fmt.Sprintf("RightScriptAttachment Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("RightScriptAttachment Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("RightScriptAttachment HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("RightScriptAttachment Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleRightScriptAttachmentRead(d, meta)
	return nil
}
func resourceRightScaleRightScriptAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRightScriptAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleRightScriptAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRoute() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names route
		// Pattern:/api/routes Vars:
		// Pattern:/api/route_tables/%s/routes Vars:route_table_id

		Create: resourceRightScaleRouteCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleRouteDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleRouteRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names route
		// Pattern:/api/routes/%s Vars:id
		// Pattern:/api/route_tables/%s/routes/%s Vars:route_table_id,id

		Update: resourceRightScaleRouteUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key route. -- Submatch: ["route[destination_cidr_block]" "route" "destination_cidr_block"] at idx 0 -- cmdFlagName: route[destination_cidr_block] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[route_table_href]" "route" "route_table_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[next_hop_href]" "route" "next_hop_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[next_hop_type]" "route" "next_hop_type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[description]" "route" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[next_hop_ip]" "route" "next_hop_ip"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[destination_cidr_block]" "route" "destination_cidr_block"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[next_hop_href]" "route" "next_hop_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[next_hop_type]" "route" "next_hop_type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[description]" "route" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			// Discovered at key route of parent map and updating -- Submatch: ["route[next_hop_ip]" "route" "next_hop_ip"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route because temp_attr wasn't null, which it never will be
			"route": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"destination_cidr_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"next_hop_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"next_hop_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleRouteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Route", "create", "/api/routes", params)
	if err != nil {
		message := fmt.Sprintf("Route Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Route Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Route HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Route Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleRouteRead(d, meta)
	return nil
}
func resourceRightScaleRouteDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRouteRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRouteTable() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names route_table
		// Pattern:/api/route_tables Vars:

		Create: resourceRightScaleRouteTableCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleRouteTableDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleRouteTableRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names route_table
		// Pattern:/api/route_tables/%s Vars:id

		Update: resourceRightScaleRouteTableUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key route_table. -- Submatch: ["route_table[network_href]" "route_table" "network_href"] at idx 0 -- cmdFlagName: route_table[network_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route_table because temp_attr wasn't null, which it never will be
			// Discovered at key route_table of parent map and updating -- Submatch: ["route_table[description]" "route_table" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route_table because temp_attr wasn't null, which it never will be
			// Discovered at key route_table of parent map and updating -- Submatch: ["route_table[cloud_href]" "route_table" "cloud_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route_table because temp_attr wasn't null, which it never will be
			// Discovered at key route_table of parent map and updating -- Submatch: ["route_table[name]" "route_table" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route_table because temp_attr wasn't null, which it never will be
			// Discovered at key route_table of parent map and updating -- Submatch: ["route_table[description]" "route_table" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route_table because temp_attr wasn't null, which it never will be
			// Discovered at key route_table of parent map and updating -- Submatch: ["route_table[name]" "route_table" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at route_table because temp_attr wasn't null, which it never will be
			"route_table": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"routes": &schema.Schema{
				Type:     schema.TypeList, //[]Route,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleRouteTableCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("RouteTable", "create", "/api/route_tables", params)
	if err != nil {
		message := fmt.Sprintf("RouteTable Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("RouteTable Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("RouteTable HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("RouteTable Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleRouteTableRead(d, meta)
	return nil
}
func resourceRightScaleRouteTableDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRouteTableRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleRouteTableUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRunnableBinding() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names runnable_binding
		// Pattern:/api/server_templates/%s/runnable_bindings Vars:server_template_id

		Create: resourceRightScaleRunnableBindingCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleRunnableBindingDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		// ACTION: multi_update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names runnable_bindings

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleRunnableBindingRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"recipe": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"sequence": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key runnable_binding. -- Submatch: ["runnable_binding[right_script_href]" "runnable_binding" "right_script_href"] at idx 0 -- cmdFlagName: runnable_binding[right_script_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at runnable_binding because temp_attr wasn't null, which it never will be
			// Discovered at key runnable_binding of parent map and updating -- Submatch: ["runnable_binding[position]" "runnable_binding" "position"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at runnable_binding because temp_attr wasn't null, which it never will be
			// Discovered at key runnable_binding of parent map and updating -- Submatch: ["runnable_binding[sequence]" "runnable_binding" "sequence"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at runnable_binding because temp_attr wasn't null, which it never will be
			// Discovered at key runnable_binding of parent map and updating -- Submatch: ["runnable_binding[recipe]" "runnable_binding" "recipe"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at runnable_binding because temp_attr wasn't null, which it never will be
			"runnable_binding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"position": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleRunnableBindingCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("RunnableBinding", "create", "/api/server_templates/%s/runnable_bindings", params)
	if err != nil {
		message := fmt.Sprintf("RunnableBinding Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("RunnableBinding Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("RunnableBinding HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("RunnableBinding Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleRunnableBindingRead(d, meta)
	return nil
}
func resourceRightScaleRunnableBindingDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleRunnableBindingRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleScheduler() *schema.Resource {
	return &schema.Resource{
		// ACTION: schedule_recipe
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names arguments,audit_id,audit_period,formal_values,policy,recipe,recipe_id,thread

		// ACTION: schedule_right_script
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names arguments,audit_id,audit_period,formal_values,policy,right_script,right_script_id,thread

		Schema: map[string]*schema.Schema{},
	}
}

func resourceRightScaleSecurityGroup() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names security_group
		// Pattern:/api/clouds/%s/security_groups Vars:cloud_id

		Create: resourceRightScaleSecurityGroupCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleSecurityGroupDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleSecurityGroupRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key security_group. -- Submatch: ["security_group[network_href]" "security_group" "network_href"] at idx 0 -- cmdFlagName: security_group[network_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group because temp_attr wasn't null, which it never will be
			// Discovered at key security_group of parent map and updating -- Submatch: ["security_group[description]" "security_group" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group because temp_attr wasn't null, which it never will be
			// Discovered at key security_group of parent map and updating -- Submatch: ["security_group[name]" "security_group" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group because temp_attr wasn't null, which it never will be
			"security_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"href": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleSecurityGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("SecurityGroup", "create", "/api/clouds/%s/security_groups", params)
	if err != nil {
		message := fmt.Sprintf("SecurityGroup Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("SecurityGroup Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("SecurityGroup HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("SecurityGroup Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleSecurityGroupRead(d, meta)
	return nil
}
func resourceRightScaleSecurityGroupDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleSecurityGroupRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleSecurityGroupRule() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names security_group_rule
		// Pattern:/api/security_group_rules Vars:
		// Pattern:/api/clouds/%s/security_groups/%s/security_group_rules Vars:cloud_id,security_group_id

		Create: resourceRightScaleSecurityGroupRuleCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleSecurityGroupRuleDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleSecurityGroupRuleRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names security_group_rule
		// Pattern:/api/security_group_rules/%s Vars:id
		// Pattern:/api/clouds/%s/security_groups/%s/security_group_rules/%s Vars:cloud_id,security_group_id,id

		Update: resourceRightScaleSecurityGroupRuleUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"cidr_ips": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			// DEBUG INFO
			"end_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"icmp_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"source_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"group_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created as a child node when an existing child of the same key was not found. Parent was security_group_rule -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[security_group_href]" "security_group_rule" "security_group_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group_rule because temp_attr wasn't null, which it never will be
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[group_owner]" "security_group_rule" "group_owner"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group_rule because temp_attr wasn't null, which it never will be
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[source_type]" "security_group_rule" "source_type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group_rule because temp_attr wasn't null, which it never will be
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[group_name]" "security_group_rule" "group_name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group_rule because temp_attr wasn't null, which it never will be
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[direction]" "security_group_rule" "direction"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group_rule because temp_attr wasn't null, which it never will be
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[protocol]" "security_group_rule" "protocol"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group_rule because temp_attr wasn't null, which it never will be
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[cidr_ips]" "security_group_rule" "cidr_ips"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group_rule because temp_attr wasn't null, which it never will be
			// Discovered at key security_group_rule of parent map and updating -- Submatch: ["security_group_rule[description]" "security_group_rule" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at security_group_rule because temp_attr wasn't null, which it never will be
			"security_group_rule": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// DEBUG INFO Child at parent[security_group_rule][protocol_details] -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] at idx 0 -- cmdFlagName: security_group_rule[protocol_details][start_port]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] -- CmdFlag: security_group_rule[protocol_details][start_port]
						// Discovered as a child of security_group_rule at key protocol_details
						// My children became the parent map after being discovered as a child of security_group_rule at key protocol_details -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] -- CmdFlag: security_group_rule[protocol_details][icmp_type]
						// (At Root) My children became the parent map. -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] -- CmdFlag: security_group_rule[protocol_details][icmp_type]
						// Discovered as a child of security_group_rule at key protocol_details
						// My children became the parent map after being discovered as a child of security_group_rule at key protocol_details -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] -- CmdFlag: security_group_rule[protocol_details][icmp_code]
						// (At Root) My children became the parent map. -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] -- CmdFlag: security_group_rule[protocol_details][icmp_code]
						// Discovered as a child of security_group_rule at key protocol_details
						// My children became the parent map after being discovered as a child of security_group_rule at key protocol_details -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] -- CmdFlag: security_group_rule[protocol_details][end_port]
						// (At Root) My children became the parent map. -- Submatch: ["security_group_rule[protocol_details]" "security_group_rule" "protocol_details"] -- CmdFlag: security_group_rule[protocol_details][end_port]
						"protocol_details": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// DEBUG INFO Created for key end_port. -- Submatch: ["[end_port]" "" "end_port"] at idx 1 -- cmdFlagName: security_group_rule[protocol_details][end_port] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at end_port because temp_attr wasn't null, which it never will be
									"end_port": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key icmp_code. -- Submatch: ["[icmp_code]" "" "icmp_code"] at idx 1 -- cmdFlagName: security_group_rule[protocol_details][icmp_code] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at icmp_code because temp_attr wasn't null, which it never will be
									"icmp_code": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key icmp_type. -- Submatch: ["[icmp_type]" "" "icmp_type"] at idx 1 -- cmdFlagName: security_group_rule[protocol_details][icmp_type] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at icmp_type because temp_attr wasn't null, which it never will be
									"icmp_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key start_port. -- Submatch: ["[start_port]" "" "start_port"] at idx 1 -- cmdFlagName: security_group_rule[protocol_details][start_port] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at start_port because temp_attr wasn't null, which it never will be
									"start_port": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"icmp_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"group_owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"href": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"start_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleSecurityGroupRuleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("SecurityGroupRule", "create", "/api/security_group_rules", params)
	if err != nil {
		message := fmt.Sprintf("SecurityGroupRule Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("SecurityGroupRule Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("SecurityGroupRule HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("SecurityGroupRule Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleSecurityGroupRuleRead(d, meta)
	return nil
}
func resourceRightScaleSecurityGroupRuleDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleSecurityGroupRuleRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleSecurityGroupRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServer() *schema.Resource {
	return &schema.Resource{
		// ACTION: clone
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names server
		// Pattern:/api/servers Vars:
		// Pattern:/api/deployments/%s/servers Vars:deployment_id

		Create: resourceRightScaleServerCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleServerDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: launch
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names api_behavior,count,inputs

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleServerRead,

		// ACTION: terminate
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: unwrap
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names server
		// Pattern:/api/servers/%s Vars:id
		// Pattern:/api/deployments/%s/servers/%s Vars:deployment_id,id

		Update: resourceRightScaleServerUpdate,

		// ACTION: wrap_instance
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names server

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created as a child node when an existing child of the same key was not found. Parent was server -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server of parent map and updating -- Submatch: ["server[deployment_href]" "server" "deployment_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server because temp_attr wasn't null, which it never will be
			// Discovered at key server of parent map and updating -- Submatch: ["server[description]" "server" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server because temp_attr wasn't null, which it never will be
			// Discovered at key server of parent map and updating -- Submatch: ["server[optimized]" "server" "optimized"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server because temp_attr wasn't null, which it never will be
			// Discovered at key server of parent map and updating -- Submatch: ["server[name]" "server" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server because temp_attr wasn't null, which it never will be
			// Discovered at key server of parent map and updating -- Submatch: ["server[automatic_instance_store_mapping]" "server" "automatic_instance_store_mapping"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server because temp_attr wasn't null, which it never will be
			// Discovered at key server of parent map and updating -- Submatch: ["server[root_volume_size]" "server" "root_volume_size"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server because temp_attr wasn't null, which it never will be
			// Discovered at key server of parent map and updating -- Submatch: ["server[description]" "server" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server because temp_attr wasn't null, which it never will be
			// Discovered at key server of parent map and updating -- Submatch: ["server[optimized]" "server" "optimized"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server because temp_attr wasn't null, which it never will be
			// Discovered at key server of parent map and updating -- Submatch: ["server[name]" "server" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server because temp_attr wasn't null, which it never will be
			"server": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// DEBUG INFO Child at parent[server][instance] -- Submatch: ["server[instance]" "server" "instance"] at idx 0 -- cmdFlagName: server[instance][cloud_specific_attributes][automatic_instance_store_mapping]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][automatic_instance_store_mapping]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][root_volume_performance]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][root_volume_performance]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][iam_instance_profile]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][iam_instance_profile]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][root_volume_type_uid]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][root_volume_type_uid]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][root_volume_size]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][root_volume_size]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][memory_mb]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][memory_mb]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][num_cores]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][num_cores]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][disk_gb]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_specific_attributes][disk_gb]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][associate_public_ip_address]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][associate_public_ip_address]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][multi_cloud_image_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][multi_cloud_image_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][ip_forwarding_enabled]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][ip_forwarding_enabled]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][security_group_hrefs][]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][security_group_hrefs][]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][placement_group_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][placement_group_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][server_template_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][server_template_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][instance_type_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][instance_type_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][ramdisk_image_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][ramdisk_image_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][kernel_image_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][kernel_image_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][datacenter_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][datacenter_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][inputs][][value]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][inputs][][value]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][inputs]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][inputs]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][inputs][][name]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][inputs][][name]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][subnet_hrefs][]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][subnet_hrefs][]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][ssh_key_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][ssh_key_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][image_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][image_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_href]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][cloud_href]
						// Discovered as a child of server at key instance
						// My children became the parent map after being discovered as a child of server at key instance -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][user_data]
						// (At Root) My children became the parent map. -- Submatch: ["server[instance]" "server" "instance"] -- CmdFlag: server[instance][user_data]
						"instance": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// DEBUG INFO Created for key associate_public_ip_address. -- Submatch: ["[associate_public_ip_address]" "" "associate_public_ip_address"] at idx 1 -- cmdFlagName: server[instance][associate_public_ip_address] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at associate_public_ip_address because temp_attr wasn't null, which it never will be
									"associate_public_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key cloud_href. -- Submatch: ["[cloud_href]" "" "cloud_href"] at idx 1 -- cmdFlagName: server[instance][cloud_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at cloud_href because temp_attr wasn't null, which it never will be
									"cloud_href": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									// DEBUG INFO Created for key datacenter_href. -- Submatch: ["[datacenter_href]" "" "datacenter_href"] at idx 1 -- cmdFlagName: server[instance][datacenter_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at datacenter_href because temp_attr wasn't null, which it never will be
									"datacenter_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key image_href. -- Submatch: ["[image_href]" "" "image_href"] at idx 1 -- cmdFlagName: server[instance][image_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at image_href because temp_attr wasn't null, which it never will be
									"image_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created as a child node when an existing child of the same key was not found. Parent was inputs -- Submatch: ["[inputs]" "" "inputs"] at idx 1 -- Operating at root: false
									// (Below Root) My children became the parent map. -- Submatch: ["[inputs]" "" "inputs"] -- CmdFlag: server[instance][inputs][][name]
									"inputs": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// DEBUG INFO Child at parent[inputs][inputs] -- Submatch: ["[inputs]" "" "inputs"] at idx 1 -- cmdFlagName: server[instance][inputs][][name]
												"inputs": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												// DEBUG INFO Created for key name. -- Submatch: ["[name]" "" "name"] at idx 3 -- cmdFlagName: server[instance][inputs][][name] -- Operating at root: false
												// Matched at subMatchIdx == last -- Operating at root: false
												// Assigned to parent map at name because temp_attr wasn't null, which it never will be
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									// DEBUG INFO Created for key instance_type_href. -- Submatch: ["[instance_type_href]" "" "instance_type_href"] at idx 1 -- cmdFlagName: server[instance][instance_type_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at instance_type_href because temp_attr wasn't null, which it never will be
									"instance_type_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key ip_forwarding_enabled. -- Submatch: ["[ip_forwarding_enabled]" "" "ip_forwarding_enabled"] at idx 1 -- cmdFlagName: server[instance][ip_forwarding_enabled] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at ip_forwarding_enabled because temp_attr wasn't null, which it never will be
									"ip_forwarding_enabled": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key kernel_image_href. -- Submatch: ["[kernel_image_href]" "" "kernel_image_href"] at idx 1 -- cmdFlagName: server[instance][kernel_image_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at kernel_image_href because temp_attr wasn't null, which it never will be
									"kernel_image_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key multi_cloud_image_href. -- Submatch: ["[multi_cloud_image_href]" "" "multi_cloud_image_href"] at idx 1 -- cmdFlagName: server[instance][multi_cloud_image_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at multi_cloud_image_href because temp_attr wasn't null, which it never will be
									"multi_cloud_image_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key placement_group_href. -- Submatch: ["[placement_group_href]" "" "placement_group_href"] at idx 1 -- cmdFlagName: server[instance][placement_group_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at placement_group_href because temp_attr wasn't null, which it never will be
									"placement_group_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key ramdisk_image_href. -- Submatch: ["[ramdisk_image_href]" "" "ramdisk_image_href"] at idx 1 -- cmdFlagName: server[instance][ramdisk_image_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at ramdisk_image_href because temp_attr wasn't null, which it never will be
									"ramdisk_image_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key server_template_href. -- Submatch: ["[server_template_href]" "" "server_template_href"] at idx 1 -- cmdFlagName: server[instance][server_template_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at server_template_href because temp_attr wasn't null, which it never will be
									"server_template_href": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									// DEBUG INFO Created for key ssh_key_href. -- Submatch: ["[ssh_key_href]" "" "ssh_key_href"] at idx 1 -- cmdFlagName: server[instance][ssh_key_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at ssh_key_href because temp_attr wasn't null, which it never will be
									"ssh_key_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key user_data. -- Submatch: ["[user_data]" "" "user_data"] at idx 1 -- cmdFlagName: server[instance][user_data] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at user_data because temp_attr wasn't null, which it never will be
									"user_data": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"optimized": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleServerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Server", "create", "/api/servers", params)
	if err != nil {
		message := fmt.Sprintf("Server Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Server Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Server HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Server Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleServerRead(d, meta)
	return nil
}
func resourceRightScaleServerDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServerRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServerUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServerArray() *schema.Resource {
	return &schema.Resource{
		// ACTION: clone
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names server_array
		// Pattern:/api/server_arrays Vars:
		// Pattern:/api/deployments/%s/server_arrays Vars:deployment_id

		Create: resourceRightScaleServerArrayCreate,

		// ACTION: current_instances
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleServerArrayDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: launch
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names api_behavior,count,inputs

		// ACTION: multi_run_executable
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names ignore_lock,inputs,recipe_name,right_script_href

		// ACTION: multi_terminate
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names terminate_all

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleServerArrayRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names server_array
		// Pattern:/api/server_arrays/%s Vars:id
		// Pattern:/api/deployments/%s/server_arrays/%s Vars:deployment_id,id

		Update: resourceRightScaleServerArrayUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"array_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"optimized": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created as a child node when an existing child of the same key was not found. Parent was server_array -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[deployment_href]" "server_array" "deployment_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[description]" "server_array" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[array_type]" "server_array" "array_type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[optimized]" "server_array" "optimized"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[state]" "server_array" "state"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[name]" "server_array" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] at idx 0 -- Operating at root: true
			// This is a top level attribute.
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[deployment_href]" "server_array" "deployment_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[description]" "server_array" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[array_type]" "server_array" "array_type"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[optimized]" "server_array" "optimized"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[state]" "server_array" "state"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			// Discovered at key server_array of parent map and updating -- Submatch: ["server_array[name]" "server_array" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_array because temp_attr wasn't null, which it never will be
			"server_array": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// DEBUG INFO Child at parent[server_array][datacenter_policy] -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] at idx 0 -- cmdFlagName: server_array[datacenter_policy][][datacenter_href]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][datacenter_href]
						// Discovered as a child of server_array at key datacenter_policy
						// My children became the parent map after being discovered as a child of server_array at key datacenter_policy -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][weight]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][weight]
						// Discovered as a child of server_array at key datacenter_policy
						// My children became the parent map after being discovered as a child of server_array at key datacenter_policy -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][max]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][max]
						// Discovered as a child of server_array at key datacenter_policy
						// My children became the parent map after being discovered as a child of server_array at key datacenter_policy -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][datacenter_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][datacenter_href]
						// Discovered as a child of server_array at key datacenter_policy
						// My children became the parent map after being discovered as a child of server_array at key datacenter_policy -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][weight]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][weight]
						// Discovered as a child of server_array at key datacenter_policy
						// My children became the parent map after being discovered as a child of server_array at key datacenter_policy -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][max]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[datacenter_policy]" "server_array" "datacenter_policy"] -- CmdFlag: server_array[datacenter_policy][][max]
						"datacenter_policy": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// DEBUG INFO Created for key datacenter_href. -- Submatch: ["[datacenter_href]" "" "datacenter_href"] at idx 2 -- cmdFlagName: server_array[datacenter_policy][][datacenter_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at datacenter_href because temp_attr wasn't null, which it never will be
									// Discovered at key datacenter_href of parent map and updating -- Submatch: ["[datacenter_href]" "" "datacenter_href"] at idx 2 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at datacenter_href because temp_attr wasn't null, which it never will be
									"datacenter_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key max. -- Submatch: ["[max]" "" "max"] at idx 2 -- cmdFlagName: server_array[datacenter_policy][][max] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at max because temp_attr wasn't null, which it never will be
									// Discovered at key max of parent map and updating -- Submatch: ["[max]" "" "max"] at idx 2 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at max because temp_attr wasn't null, which it never will be
									"max": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key weight. -- Submatch: ["[weight]" "" "weight"] at idx 2 -- cmdFlagName: server_array[datacenter_policy][][weight] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at weight because temp_attr wasn't null, which it never will be
									// Discovered at key weight of parent map and updating -- Submatch: ["[weight]" "" "weight"] at idx 2 -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at weight because temp_attr wasn't null, which it never will be
									"weight": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						// DEBUG INFO Child at parent[server_array][elasticity_params] -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] at idx 0 -- cmdFlagName: server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][collect_audit_entries]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][collect_audit_entries]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][alert_specific_params][voters_tag_predicate]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][alert_specific_params][voters_tag_predicate]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][algorithm]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][algorithm]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][alert_specific_params][decision_threshold]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][alert_specific_params][decision_threshold]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][max_age]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][max_age]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][regexp]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][regexp]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_calm_time]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_calm_time]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_down_by]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_down_by]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][min_count]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][min_count]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][max_count]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][max_count]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_up_by]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_up_by]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][bounds][max_count]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][bounds][max_count]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][bounds][min_count]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][bounds][min_count]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][time]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][time]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][day]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][day]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][queue_size][items_per_instance]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][collect_audit_entries]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][collect_audit_entries]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][alert_specific_params][voters_tag_predicate]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][alert_specific_params][voters_tag_predicate]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][algorithm]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][algorithm]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][alert_specific_params][decision_threshold]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][alert_specific_params][decision_threshold]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][max_age]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][max_age]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][regexp]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][queue_specific_params][item_age][regexp]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_calm_time]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_calm_time]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_down_by]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_down_by]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][max_count]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][max_count]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][min_count]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][min_count]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_up_by]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][pacing][resize_up_by]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][bounds][min_count]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][bounds][min_count]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][bounds][max_count]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][bounds][max_count]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][time]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][time]
						// Discovered as a child of server_array at key elasticity_params
						// My children became the parent map after being discovered as a child of server_array at key elasticity_params -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][day]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[elasticity_params]" "server_array" "elasticity_params"] -- CmdFlag: server_array[elasticity_params][schedule][][day]
						"elasticity_params": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
						},
						// DEBUG INFO Child at parent[server_array][instance] -- Submatch: ["server_array[instance]" "server_array" "instance"] at idx 0 -- cmdFlagName: server_array[instance][cloud_specific_attributes][automatic_instance_store_mapping]
						// My children are being initialized, they were []
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_specific_attributes][automatic_instance_store_mapping]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_specific_attributes][root_volume_performance]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_specific_attributes][root_volume_performance]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_specific_attributes][root_volume_type_uid]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_specific_attributes][root_volume_type_uid]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_specific_attributes][iam_instance_profile]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_specific_attributes][iam_instance_profile]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_specific_attributes][root_volume_size]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_specific_attributes][root_volume_size]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][associate_public_ip_address]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][associate_public_ip_address]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][multi_cloud_image_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][multi_cloud_image_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][ip_forwarding_enabled]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][ip_forwarding_enabled]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][placement_group_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][placement_group_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][security_group_hrefs][]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][security_group_hrefs][]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][server_template_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][server_template_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][ramdisk_image_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][ramdisk_image_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][instance_type_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][instance_type_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][kernel_image_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][kernel_image_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][datacenter_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][datacenter_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][inputs][][value]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][inputs][][value]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][inputs]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][inputs]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][inputs][][name]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][inputs][][name]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][subnet_hrefs][]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][subnet_hrefs][]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][ssh_key_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][ssh_key_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][cloud_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][image_href]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][image_href]
						// Discovered as a child of server_array at key instance
						// My children became the parent map after being discovered as a child of server_array at key instance -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][user_data]
						// (At Root) My children became the parent map. -- Submatch: ["server_array[instance]" "server_array" "instance"] -- CmdFlag: server_array[instance][user_data]
						"instance": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// DEBUG INFO Created for key associate_public_ip_address. -- Submatch: ["[associate_public_ip_address]" "" "associate_public_ip_address"] at idx 1 -- cmdFlagName: server_array[instance][associate_public_ip_address] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at associate_public_ip_address because temp_attr wasn't null, which it never will be
									"associate_public_ip_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key cloud_href. -- Submatch: ["[cloud_href]" "" "cloud_href"] at idx 1 -- cmdFlagName: server_array[instance][cloud_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at cloud_href because temp_attr wasn't null, which it never will be
									"cloud_href": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									// DEBUG INFO Created for key datacenter_href. -- Submatch: ["[datacenter_href]" "" "datacenter_href"] at idx 1 -- cmdFlagName: server_array[instance][datacenter_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at datacenter_href because temp_attr wasn't null, which it never will be
									"datacenter_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key image_href. -- Submatch: ["[image_href]" "" "image_href"] at idx 1 -- cmdFlagName: server_array[instance][image_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at image_href because temp_attr wasn't null, which it never will be
									"image_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created as a child node when an existing child of the same key was not found. Parent was inputs -- Submatch: ["[inputs]" "" "inputs"] at idx 1 -- Operating at root: false
									// (Below Root) My children became the parent map. -- Submatch: ["[inputs]" "" "inputs"] -- CmdFlag: server_array[instance][inputs][][name]
									"inputs": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// DEBUG INFO Child at parent[inputs][inputs] -- Submatch: ["[inputs]" "" "inputs"] at idx 1 -- cmdFlagName: server_array[instance][inputs][][name]
												"inputs": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
												// DEBUG INFO Created for key name. -- Submatch: ["[name]" "" "name"] at idx 3 -- cmdFlagName: server_array[instance][inputs][][name] -- Operating at root: false
												// Matched at subMatchIdx == last -- Operating at root: false
												// Assigned to parent map at name because temp_attr wasn't null, which it never will be
												"name": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									// DEBUG INFO Created for key instance_type_href. -- Submatch: ["[instance_type_href]" "" "instance_type_href"] at idx 1 -- cmdFlagName: server_array[instance][instance_type_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at instance_type_href because temp_attr wasn't null, which it never will be
									"instance_type_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key ip_forwarding_enabled. -- Submatch: ["[ip_forwarding_enabled]" "" "ip_forwarding_enabled"] at idx 1 -- cmdFlagName: server_array[instance][ip_forwarding_enabled] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at ip_forwarding_enabled because temp_attr wasn't null, which it never will be
									"ip_forwarding_enabled": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key kernel_image_href. -- Submatch: ["[kernel_image_href]" "" "kernel_image_href"] at idx 1 -- cmdFlagName: server_array[instance][kernel_image_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at kernel_image_href because temp_attr wasn't null, which it never will be
									"kernel_image_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key multi_cloud_image_href. -- Submatch: ["[multi_cloud_image_href]" "" "multi_cloud_image_href"] at idx 1 -- cmdFlagName: server_array[instance][multi_cloud_image_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at multi_cloud_image_href because temp_attr wasn't null, which it never will be
									"multi_cloud_image_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key placement_group_href. -- Submatch: ["[placement_group_href]" "" "placement_group_href"] at idx 1 -- cmdFlagName: server_array[instance][placement_group_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at placement_group_href because temp_attr wasn't null, which it never will be
									"placement_group_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key ramdisk_image_href. -- Submatch: ["[ramdisk_image_href]" "" "ramdisk_image_href"] at idx 1 -- cmdFlagName: server_array[instance][ramdisk_image_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at ramdisk_image_href because temp_attr wasn't null, which it never will be
									"ramdisk_image_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key server_template_href. -- Submatch: ["[server_template_href]" "" "server_template_href"] at idx 1 -- cmdFlagName: server_array[instance][server_template_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at server_template_href because temp_attr wasn't null, which it never will be
									"server_template_href": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									// DEBUG INFO Created for key ssh_key_href. -- Submatch: ["[ssh_key_href]" "" "ssh_key_href"] at idx 1 -- cmdFlagName: server_array[instance][ssh_key_href] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at ssh_key_href because temp_attr wasn't null, which it never will be
									"ssh_key_href": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									// DEBUG INFO Created for key user_data. -- Submatch: ["[user_data]" "" "user_data"] at idx 1 -- cmdFlagName: server_array[instance][user_data] -- Operating at root: false
									// Matched at subMatchIdx == last -- Operating at root: false
									// Assigned to parent map at user_data because temp_attr wasn't null, which it never will be
									"user_data": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			// DEBUG INFO
			"current_instances": &schema.Schema{
				Type:     schema.TypeList, //[]Instance,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"instances_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleServerArrayCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("ServerArray", "create", "/api/server_arrays", params)
	if err != nil {
		message := fmt.Sprintf("ServerArray Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("ServerArray Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("ServerArray HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("ServerArray Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleServerArrayRead(d, meta)
	return nil
}

func resourceRightScaleServerArrayDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServerArrayRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleServerArrayUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServerTemplate() *schema.Resource {
	return &schema.Resource{
		// ACTION: clone
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names server_template

		// ACTION: commit
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names commit_head_dependencies,commit_message,freeze_repositories

		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names server_template
		// Pattern:/api/server_templates Vars:

		Create: resourceRightScaleServerTemplateCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleServerTemplateDelete,

		// ACTION: detect_changes_in_head
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: publish
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names account_group_hrefs,allow_comments,categories,descriptions,email_comments

		// ACTION: resolve
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleServerTemplateRead,

		// ACTION: swap_repository
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names source_repository_href,target_repository_href

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names server_template
		// Pattern:/api/server_templates/%s Vars:id

		Update: resourceRightScaleServerTemplateUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"revision": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key server_template. -- Submatch: ["server_template[description]" "server_template" "description"] at idx 0 -- cmdFlagName: server_template[description] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_template because temp_attr wasn't null, which it never will be
			// Discovered at key server_template of parent map and updating -- Submatch: ["server_template[name]" "server_template" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_template because temp_attr wasn't null, which it never will be
			// Discovered at key server_template of parent map and updating -- Submatch: ["server_template[description]" "server_template" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_template because temp_attr wasn't null, which it never will be
			// Discovered at key server_template of parent map and updating -- Submatch: ["server_template[name]" "server_template" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_template because temp_attr wasn't null, which it never will be
			"server_template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"inputs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"lineage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleServerTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("ServerTemplate", "create", "/api/server_templates", params)
	if err != nil {
		message := fmt.Sprintf("ServerTemplate Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("ServerTemplate Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("ServerTemplate HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("ServerTemplate Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleServerTemplateRead(d, meta)
	return nil
}
func resourceRightScaleServerTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServerTemplateRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServerTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServerTemplateMultiCloudImage() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names server_template_multi_cloud_image
		// Pattern:/api/server_template_multi_cloud_images Vars:

		Create: resourceRightScaleServerTemplateMultiCloudImageCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleServerTemplateMultiCloudImageDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: make_default
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleServerTemplateMultiCloudImageRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key server_template_multi_cloud_image. -- Submatch: ["server_template_multi_cloud_image[multi_cloud_image_href]" "server_template_multi_cloud_image" "multi_cloud_image_href"] at idx 0 -- cmdFlagName: server_template_multi_cloud_image[multi_cloud_image_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_template_multi_cloud_image because temp_attr wasn't null, which it never will be
			// Discovered at key server_template_multi_cloud_image of parent map and updating -- Submatch: ["server_template_multi_cloud_image[server_template_href]" "server_template_multi_cloud_image" "server_template_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at server_template_multi_cloud_image because temp_attr wasn't null, which it never will be
			"server_template_multi_cloud_image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleServerTemplateMultiCloudImageCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("ServerTemplateMultiCloudImage", "create", "/api/server_template_multi_cloud_images", params)
	if err != nil {
		message := fmt.Sprintf("ServerTemplateMultiCloudImage Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("ServerTemplateMultiCloudImage Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("ServerTemplateMultiCloudImage HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("ServerTemplateMultiCloudImage Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleServerTemplateMultiCloudImageRead(d, meta)
	return nil
}
func resourceRightScaleServerTemplateMultiCloudImageDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleServerTemplateMultiCloudImageRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleSession() *schema.Resource {
	return &schema.Resource{
		// ACTION: accounts
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names email,password

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		// ACTION: index_instance_session
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleSshKey() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names ssh_key
		// Pattern:/api/clouds/%s/ssh_keys Vars:cloud_id

		Create: resourceRightScaleSshKeyCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleSshKeyDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleSshKeyRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO Created for key ssh_key. -- Submatch: ["ssh_key[name]" "ssh_key" "name"] at idx 0 -- cmdFlagName: ssh_key[name] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at ssh_key because temp_attr wasn't null, which it never will be
			"ssh_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"material": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleSshKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("SshKey", "create", "/api/clouds/%s/ssh_keys", params)
	if err != nil {
		message := fmt.Sprintf("SshKey Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("SshKey Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("SshKey HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("SshKey Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleSshKeyRead(d, meta)
	return nil
}
func resourceRightScaleSshKeyDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleSshKeyRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleSubnet() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names subnet
		// Pattern:/api/clouds/%s/instances/%s/subnets Vars:cloud_id,instance_id
		// Pattern:/api/clouds/%s/subnets Vars:cloud_id

		Create: resourceRightScaleSubnetCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleSubnetDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleSubnetRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names subnet
		// Pattern:/api/clouds/%s/instances/%s/subnets/%s Vars:cloud_id,instance_id,id
		// Pattern:/api/clouds/%s/subnets/%s Vars:cloud_id,id

		Update: resourceRightScaleSubnetUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"is_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key subnet. -- Submatch: ["subnet[datacenter_href]" "subnet" "datacenter_href"] at idx 0 -- cmdFlagName: subnet[datacenter_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at subnet because temp_attr wasn't null, which it never will be
			// Discovered at key subnet of parent map and updating -- Submatch: ["subnet[network_href]" "subnet" "network_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at subnet because temp_attr wasn't null, which it never will be
			// Discovered at key subnet of parent map and updating -- Submatch: ["subnet[description]" "subnet" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at subnet because temp_attr wasn't null, which it never will be
			// Discovered at key subnet of parent map and updating -- Submatch: ["subnet[cidr_block]" "subnet" "cidr_block"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at subnet because temp_attr wasn't null, which it never will be
			// Discovered at key subnet of parent map and updating -- Submatch: ["subnet[name]" "subnet" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at subnet because temp_attr wasn't null, which it never will be
			// Discovered at key subnet of parent map and updating -- Submatch: ["subnet[route_table_href]" "subnet" "route_table_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at subnet because temp_attr wasn't null, which it never will be
			// Discovered at key subnet of parent map and updating -- Submatch: ["subnet[description]" "subnet" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at subnet because temp_attr wasn't null, which it never will be
			// Discovered at key subnet of parent map and updating -- Submatch: ["subnet[name]" "subnet" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at subnet because temp_attr wasn't null, which it never will be
			"subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"cidr_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleSubnetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Subnet", "create", "/api/clouds/%s/instances/%s/subnets", params)
	if err != nil {
		message := fmt.Sprintf("Subnet Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Subnet Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Subnet HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Subnet Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleSubnetRead(d, meta)
	return nil
}
func resourceRightScaleSubnetDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleSubnetRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleSubnetUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleTag() *schema.Resource {
	return &schema.Resource{
		// ACTION: by_resource
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names resource_hrefs

		// ACTION: by_tag
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names include_tags_with_prefix,match_all,resource_type,tags,with_deleted

		// ACTION: multi_add
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names resource_hrefs,tags

		// ACTION: multi_delete
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names resource_hrefs,tags

		Schema: map[string]*schema.Schema{},
	}
}

func resourceRightScaleTask() *schema.Resource {
	return &schema.Resource{
		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleTaskRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"detail": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleTaskRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleUser() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names user
		// Pattern:/api/users Vars:

		Create: resourceRightScaleUserCreate,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[]
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleUserRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names user
		// Pattern:/api/users/%s Vars:id

		Update: resourceRightScaleUserUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"timezone_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"phone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"principal_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key user. -- Submatch: ["user[identity_provider_href]" "user" "identity_provider_href"] at idx 0 -- cmdFlagName: user[identity_provider_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[principal_uid]" "user" "principal_uid"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[timezone_name]" "user" "timezone_name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[first_name]" "user" "first_name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[last_name]" "user" "last_name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[password]" "user" "password"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[company]" "user" "company"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[phone]" "user" "phone"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[email]" "user" "email"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[identity_provider_href]" "user" "identity_provider_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[current_password]" "user" "current_password"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[current_email]" "user" "current_email"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[principal_uid]" "user" "principal_uid"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[timezone_name]" "user" "timezone_name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[new_password]" "user" "new_password"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[first_name]" "user" "first_name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[last_name]" "user" "last_name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[new_email]" "user" "new_email"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[company]" "user" "company"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			// Discovered at key user of parent map and updating -- Submatch: ["user[phone]" "user" "phone"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at user because temp_attr wasn't null, which it never will be
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"company": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("User", "create", "/api/users", params)
	if err != nil {
		message := fmt.Sprintf("User Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("User Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("User HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("User Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleUserRead(d, meta)
	return nil
}

func resourceRightScaleUserRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleUserUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleUserData() *schema.Resource {
	return &schema.Resource{
		// ACTION: show
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Read: resourceRightScaleUserDataRead,

		Schema: map[string]*schema.Schema{},
	}
}

func resourceRightScaleUserDataRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleVolume() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names volume
		// Pattern:/api/clouds/%s/volumes Vars:cloud_id

		Create: resourceRightScaleVolumeCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleVolumeDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleVolumeRead,

		// ACTION: update
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names volume
		// Pattern:/api/clouds/%s/volumes/%s Vars:cloud_id,id

		Update: resourceRightScaleVolumeUpdate,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key volume. -- Submatch: ["volume[parent_volume_snapshot_href]" "volume" "parent_volume_snapshot_href"] at idx 0 -- cmdFlagName: volume[parent_volume_snapshot_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[placement_group_href]" "volume" "placement_group_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[volume_type_href]" "volume" "volume_type_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[datacenter_href]" "volume" "datacenter_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[deployment_href]" "volume" "deployment_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[description]" "volume" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[encrypted]" "volume" "encrypted"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[name]" "volume" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[iops]" "volume" "iops"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[size]" "volume" "size"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			// Discovered at key volume of parent map and updating -- Submatch: ["volume[name]" "volume" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume because temp_attr wasn't null, which it never will be
			"volume": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"iops": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"volume_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleVolumeCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("Volume", "create", "/api/clouds/%s/volumes", params)
	if err != nil {
		message := fmt.Sprintf("Volume Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("Volume Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("Volume HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("Volume Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleVolumeRead(d, meta)
	return nil
}
func resourceRightScaleVolumeDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleVolumeRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceRightScaleVolumeUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleVolumeAttachment() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names volume_attachment
		// Pattern:/api/clouds/%s/instances/%s/volume_attachments Vars:cloud_id,instance_id
		// Pattern:/api/clouds/%s/volume_attachments Vars:cloud_id
		// Pattern:/api/clouds/%s/volumes/%s/volume_attachments Vars:cloud_id,volume_id
		// Pattern:/api/clouds/%s/volumes/%s/volume_attachment Vars:cloud_id,volume_id

		Create: resourceRightScaleVolumeAttachmentCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names force

		Delete: resourceRightScaleVolumeAttachmentDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleVolumeAttachmentRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"device_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key volume_attachment. -- Submatch: ["volume_attachment[instance_href]" "volume_attachment" "instance_href"] at idx 0 -- cmdFlagName: volume_attachment[instance_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key volume_attachment of parent map and updating -- Submatch: ["volume_attachment[volume_href]" "volume_attachment" "volume_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_attachment because temp_attr wasn't null, which it never will be
			// Discovered at key volume_attachment of parent map and updating -- Submatch: ["volume_attachment[device]" "volume_attachment" "device"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_attachment because temp_attr wasn't null, which it never will be
			"volume_attachment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceRightScaleVolumeAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("VolumeAttachment", "create", "/api/clouds/%s/instances/%s/volume_attachments", params)
	if err != nil {
		message := fmt.Sprintf("VolumeAttachment Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("VolumeAttachment Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("VolumeAttachment HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("VolumeAttachment Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleVolumeAttachmentRead(d, meta)
	return nil
}
func resourceRightScaleVolumeAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleVolumeAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleVolumeSnapshot() *schema.Resource {
	return &schema.Resource{
		// ACTION: create
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names volume_snapshot,volume_snapshot_copy
		// Pattern:/api/clouds/%s/volumes/%s/volume_snapshots Vars:cloud_id,volume_id
		// Pattern:/api/clouds/%s/volume_snapshots Vars:cloud_id

		Create: resourceRightScaleVolumeSnapshotCreate,

		// ACTION: destroy
		//
		// Path Param Names
		// Query Param Names
		// Payload Param Names

		Delete: resourceRightScaleVolumeSnapshotDelete,

		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleVolumeSnapshotRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO Created for key volume_snapshot_copy. -- Submatch: ["volume_snapshot_copy[volume_snapshot_href]" "volume_snapshot_copy" "volume_snapshot_href"] at idx 0 -- cmdFlagName: volume_snapshot_copy[volume_snapshot_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_snapshot_copy because temp_attr wasn't null, which it never will be
			// Discovered at key volume_snapshot_copy of parent map and updating -- Submatch: ["volume_snapshot_copy[description]" "volume_snapshot_copy" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_snapshot_copy because temp_attr wasn't null, which it never will be
			// Discovered at key volume_snapshot_copy of parent map and updating -- Submatch: ["volume_snapshot_copy[cloud_href]" "volume_snapshot_copy" "cloud_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_snapshot_copy because temp_attr wasn't null, which it never will be
			// Discovered at key volume_snapshot_copy of parent map and updating -- Submatch: ["volume_snapshot_copy[name]" "volume_snapshot_copy" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_snapshot_copy because temp_attr wasn't null, which it never will be
			"volume_snapshot_copy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO Created for key volume_snapshot. -- Submatch: ["volume_snapshot[parent_volume_href]" "volume_snapshot" "parent_volume_href"] at idx 0 -- cmdFlagName: volume_snapshot[parent_volume_href] -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_snapshot because temp_attr wasn't null, which it never will be
			// Discovered at key volume_snapshot of parent map and updating -- Submatch: ["volume_snapshot[deployment_href]" "volume_snapshot" "deployment_href"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_snapshot because temp_attr wasn't null, which it never will be
			// Discovered at key volume_snapshot of parent map and updating -- Submatch: ["volume_snapshot[description]" "volume_snapshot" "description"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_snapshot because temp_attr wasn't null, which it never will be
			// Discovered at key volume_snapshot of parent map and updating -- Submatch: ["volume_snapshot[name]" "volume_snapshot" "name"] at idx 0 -- Operating at root: true
			// Matched at subMatchIdx == last -- Operating at root: true
			// Assigned to parent map at volume_snapshot because temp_attr wasn't null, which it never will be
			"volume_snapshot": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"progress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleVolumeSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cm15.API)

	params := rsapi.APIParams{}
	req, err := client.BuildRequest("VolumeSnapshot", "create", "/api/clouds/%s/volumes/%s/volume_snapshots", params)
	if err != nil {
		message := fmt.Sprintf("VolumeSnapshot Could not create HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	resp, err := client.PerformRequest(req)
	if err != nil {
		message := fmt.Sprintf("VolumeSnapshot Could not execute HTTP request. Error: %s", err.Error())
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	}
	if resp.StatusCode != 201 {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			message := fmt.Sprintf("VolumeSnapshot HTTP Response was %d rather than 201 for a create request. Could not read the body of the HTTP request. Error: %s", resp.StatusCode, err.Error())
			log.Printf("[DEBUG] %s", message)
			return fmt.Errorf(message)
		}
		message := fmt.Sprintf("VolumeSnapshot Create failed. Status Code: %d Error: %s", resp.StatusCode, contents)
		log.Printf("[DEBUG] %s", message)
		return fmt.Errorf(message)
	} else {
		// This is where we do some processing on this thang!
	}

	// Set this resource id to RightScale HREF
	// d.SetId(string(resource.Href))

	// return resourceRightScaleVolumeSnapshotRead(d, meta)
	return nil
}
func resourceRightScaleVolumeSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleVolumeSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceRightScaleVolumeType() *schema.Resource {
	return &schema.Resource{
		// ACTION: index
		//
		// Path Param Names
		// Query Param Names filter[],view
		// Payload Param Names

		// ACTION: show
		//
		// Path Param Names
		// Query Param Names view
		// Payload Param Names

		Read: resourceRightScaleVolumeTypeRead,

		Schema: map[string]*schema.Schema{
			// DEBUG INFO
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"links": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"resource_uid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// DEBUG INFO
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},

			// DEBUG INFO
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRightScaleVolumeTypeRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

var ResourceList = map[string]*schema.Resource{
	"rightscale_account": resourceRightScaleAccount(),

	"rightscale_account_group": resourceRightScaleAccountGroup(),

	"rightscale_alert": resourceRightScaleAlert(),

	"rightscale_alert_spec": resourceRightScaleAlertSpec(),

	"rightscale_audit_entry": resourceRightScaleAuditEntry(),

	"rightscale_backup": resourceRightScaleBackup(),

	"rightscale_child_account": resourceRightScaleChildAccount(),

	"rightscale_cloud": resourceRightScaleCloud(),

	"rightscale_cloud_account": resourceRightScaleCloudAccount(),

	"rightscale_cookbook": resourceRightScaleCookbook(),

	"rightscale_cookbook_attachment": resourceRightScaleCookbookAttachment(),

	"rightscale_credential": resourceRightScaleCredential(),

	"rightscale_datacenter": resourceRightScaleDatacenter(),

	"rightscale_deployment": resourceRightScaleDeployment(),

	"rightscale_health_check": resourceRightScaleHealthCheck(),

	"rightscale_identity_provider": resourceRightScaleIdentityProvider(),

	"rightscale_image": resourceRightScaleImage(),

	"rightscale_input": resourceRightScaleInput(),

	"rightscale_instance": resourceRightScaleInstance(),

	"rightscale_instance_custom_lodgement": resourceRightScaleInstanceCustomLodgement(),

	"rightscale_instance_type": resourceRightScaleInstanceType(),

	"rightscale_ip_address": resourceRightScaleIpAddress(),

	"rightscale_ip_address_binding": resourceRightScaleIpAddressBinding(),

	"rightscale_monitoring_metric": resourceRightScaleMonitoringMetric(),

	"rightscale_multi_cloud_image": resourceRightScaleMultiCloudImage(),

	"rightscale_multi_cloud_image_setting": resourceRightScaleMultiCloudImageSetting(),

	"rightscale_network": resourceRightScaleNetwork(),

	"rightscale_network_gateway": resourceRightScaleNetworkGateway(),

	"rightscale_network_option_group": resourceRightScaleNetworkOptionGroup(),

	"rightscale_network_option_group_attachment": resourceRightScaleNetworkOptionGroupAttachment(),

	"rightscale_oauth2": resourceRightScaleOauth2(),

	"rightscale_permission": resourceRightScalePermission(),

	"rightscale_placement_group": resourceRightScalePlacementGroup(),

	"rightscale_preference": resourceRightScalePreference(),

	"rightscale_publication": resourceRightScalePublication(),

	"rightscale_publication_lineage": resourceRightScalePublicationLineage(),

	"rightscale_recurring_volume_attachment": resourceRightScaleRecurringVolumeAttachment(),

	"rightscale_repository": resourceRightScaleRepository(),

	"rightscale_repository_asset": resourceRightScaleRepositoryAsset(),

	"rightscale_right_script": resourceRightScaleRightScript(),

	"rightscale_right_script_attachment": resourceRightScaleRightScriptAttachment(),

	"rightscale_route": resourceRightScaleRoute(),

	"rightscale_route_table": resourceRightScaleRouteTable(),

	"rightscale_runnable_binding": resourceRightScaleRunnableBinding(),

	"rightscale_scheduler": resourceRightScaleScheduler(),

	"rightscale_security_group": resourceRightScaleSecurityGroup(),

	"rightscale_security_group_rule": resourceRightScaleSecurityGroupRule(),

	"rightscale_server": resourceRightScaleServer(),

	"rightscale_server_array": resourceRightScaleServerArray(),

	"rightscale_server_template": resourceRightScaleServerTemplate(),

	"rightscale_server_template_multi_cloud_image": resourceRightScaleServerTemplateMultiCloudImage(),

	"rightscale_session": resourceRightScaleSession(),

	"rightscale_ssh_key": resourceRightScaleSshKey(),

	"rightscale_subnet": resourceRightScaleSubnet(),

	"rightscale_tag": resourceRightScaleTag(),

	"rightscale_task": resourceRightScaleTask(),

	"rightscale_user": resourceRightScaleUser(),

	"rightscale_user_data": resourceRightScaleUserData(),

	"rightscale_volume": resourceRightScaleVolume(),

	"rightscale_volume_attachment": resourceRightScaleVolumeAttachment(),

	"rightscale_volume_snapshot": resourceRightScaleVolumeSnapshot(),

	"rightscale_volume_type": resourceRightScaleVolumeType(),
}