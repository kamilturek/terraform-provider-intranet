package intranet

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/kamilturek/intranet-go"
)

func resourceHourEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceHourEntryCreate,
		Read:   resourceHourEntryRead,
		Update: resourceHourEntryUpdate,
		Delete: resourceHourEntryDelete,

		Schema: map[string]*schema.Schema{
			"date": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"project_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"project_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ticket_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time": {
				Type:         schema.TypeFloat,
				Required:     true,
				ValidateFunc: validation.FloatAtLeast(0),
			},
		},
	}
}

func resourceHourEntryCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*intranet.Client)

	input := &intranet.CreateHourEntryInput{
		Date:        d.Get("date").(string),
		Description: d.Get("description").(string),
		ProjectID:   d.Get("project_id").(int),
		Time:        d.Get("time").(float64),
	}

	if v, ok := d.GetOk("ticket_id"); ok {
		input.TicketID = v.(string)
	}

	output, err := client.CreateHourEntry(input)
	if err != nil {
		return fmt.Errorf("error creating hour entry: %w", err)
	}

	d.SetId(output.ID)
	d.Set("date", d.Get("date"))

	return resourceHourEntryRead(d, meta)
}

func resourceHourEntryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*intranet.Client)

	input := &intranet.ListHourEntriesInput{
		Date: d.Get("date").(string),
	}

	output, err := client.ListHourEntries(input)
	if err != nil {
		return fmt.Errorf("error listing hour entries")
	}

	var entry *intranet.HourEntry

	for _, e := range output.Entries {
		if strconv.Itoa(e.ID) == d.Id() {
			entry = &e
			break
		}
	}

	if entry == nil {
		if !d.IsNewResource() {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("hour entry not found")
	}

	d.Set("description", entry.Description)
	d.Set("time", entry.Time)
	d.Set("project_id", entry.Project.ID)
	d.Set("project_name", entry.Project.Name)

	return nil
}

func resourceHourEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*intranet.Client)

	input := &intranet.UpdateHourEntryInput{
		Date:      d.Get("date").(string),
		ID:        d.Id(),
		ProjectID: d.Get("project_id").(int),
	}

	if v, ok := d.GetOk("description"); ok {
		input.Description = v.(string)
	}

	if v, ok := d.GetOk("ticket_id"); ok {
		input.TicketID = v.(string)
	}

	if v, ok := d.GetOk("time"); ok {
		input.Time = v.(float64)
	}

	_, err := client.UpdateHourEntry(input)
	if err != nil {
		return fmt.Errorf("error updating hour entry: %w", err)
	}

	return resourceHourEntryRead(d, meta)
}

func resourceHourEntryDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*intranet.Client)

	input := &intranet.DeleteHourEntryInput{
		ID: d.Id(),
	}

	err := client.DeleteHourEntry(input)
	if err != nil {
		return fmt.Errorf("error deleting hour entry: %w", err)
	}

	return nil
}
