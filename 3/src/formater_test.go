
package main

import "testing"

func TestGreeting(t *testing.T) {
	completeGreeting := greeting("Code.educationRocks")
	if completeGreeting != "<b> Code.educationRocks </b>" {
		t.Errorf("greeting is incorrect, got: %s, want: <b> Code.educationRocks </b>", completeGreeting);
	}
}
