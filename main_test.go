package main

import (
	"testing"
)

func callback(m string) {}

func TestGameUpdateSync(t *testing.T) {
	engine := game{}
	if err := engine.Init(); err != nil {
		t.Error(err)
	}
	if err := engine.Update([]byte{}, callback); err != nil {
		t.Error(err)
	}
	if state := engine.GetState(); len(state) != 0 {
		t.Error()
	}
	payload := []byte(`{"command":"system_sync","data":{"name":"p1","x":1,"y":1}}`)
	if err := engine.Update(payload, callback); err != nil {
		t.Error(err)
	}
	state := engine.GetState()
	if len(state) == 0 || string(state) != `{"command":"system_sync","data":[{"name":"p1","x":1,"y":1}]}` {
		t.Error()
	}
}

func TestGameUpdateAdd(t *testing.T) {
	engine := game{}
	if err := engine.Init(); err != nil {
		t.Error(err)
	}
	if err := engine.Update([]byte{}, callback); err != nil {
		t.Error(err)
	}
	if state := engine.GetState(); len(state) != 0 {
		t.Error()
	}
	payload := []byte(`{"command":"player_add","data":{"name":"p1","x":1,"y":1}}`)
	if err := engine.Update(payload, callback); err != nil {
		t.Error(err)
	}
	state := engine.GetState()
	if len(state) == 0 || string(state) != `{"command":"player_add","data":{"name":"p1","x":1,"y":1}}` {
		t.Error()
	}
}

func TestGameUpdateRemove(t *testing.T) {
	engine := game{}
	if err := engine.Init(); err != nil {
		t.Error(err)
	}
	if err := engine.Update([]byte{}, callback); err != nil {
		t.Error(err)
	}
	if state := engine.GetState(); len(state) != 0 {
		t.Error()
	}
	payload := []byte(`{"command":"player_remove","data":{"name":"p1","x":1,"y":1}}`)
	if err := engine.Update(payload, callback); err != nil {
		t.Error(err)
	}
	state := engine.GetState()
	if len(state) == 0 || string(state) != `{"command":"player_remove","data":{"name":"p1","x":1,"y":1}}` {
		t.Error()
	}
}

func TestGameUpdateSyncAfterAdd(t *testing.T) {
	engine := game{}
	if err := engine.Init(); err != nil {
		t.Error(err)
	}
	if err := engine.Update([]byte{}, callback); err != nil {
		t.Error(err)
	}
	if state := engine.GetState(); len(state) != 0 {
		t.Error()
	}
	payload := []byte(`{"command":"player_add","data":{"name":"p1","x":1,"y":1}}`)
	if err := engine.Update(payload, callback); err != nil {
		t.Error(err)
	}
	state := engine.GetState()
	if len(state) == 0 || string(state) != `{"command":"player_add","data":{"name":"p1","x":1,"y":1}}` {
		t.Error()
	}
	payload = []byte(`{"command":"system_sync","data":{"name":"p2","x":2.1,"y":2.1}}`)
	if err := engine.Update(payload, callback); err != nil {
		t.Error(err)
	}
	state = engine.GetState()
	if len(state) == 0 || string(state) != `{"command":"system_sync","data":[{"name":"p1","x":1,"y":1},{"name":"p2","x":2.1,"y":2.1}]}` {
		t.Error()
	}
}
