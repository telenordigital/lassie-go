package lassie

import (
	"reflect"
	"testing"
)

func TestOutput(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}

	app, err := c.CreateApplication(Application{Tags: map[string]string{"name": "Lassie test app"}})
	if err != nil {
		t.Fatal(err)
	}
	defer c.DeleteApplication(app.EUI)

	op, err := c.CreateOutput(app.EUI, &MQTTConfig{Endpoint: "localhost", Port: 1883})
	if err != nil {
		t.Fatal(err)
	}
	cfg, err := op.MQTTConfig()
	if err != nil {
		t.Fatal(err)
	}

	cfg.Username = "foo"
	cfg.Port = 1884

	updatedOp, err := c.UpdateOutput(app.EUI, op.EUI, &cfg)
	if err != nil {
		t.Fatal(err)
	}

	updatedCfg, err := updatedOp.MQTTConfig()
	if err != nil {
		t.Fatal(err)
	}
	if updatedCfg.Port != 1884 || updatedCfg.Username != "foo" {
		t.Fatal("MQTT config wasn't updated")
	}

	existing, err := c.Output(app.EUI, op.EUI)
	if err != nil {
		t.Fatal(err)
	}
	existingCfg, err := existing.MQTTConfig()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(existingCfg, updatedCfg) {
		t.Fatalf("Not the same: %+v != %+v", existingCfg, updatedCfg)
	}
	list, err := c.Outputs(app.EUI)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 || list[0].EUI != op.EUI {
		t.Fatal("List did not contain output")
	}
	if cfg, err := list[0].MQTTConfig(); err != nil {
		t.Fatalf("Error retrieving config from list: %v %v", cfg, err)
	}
	if err := c.DeleteOutput(app.EUI, op.EUI); err != nil {
		t.Fatal(err)
	}
}
