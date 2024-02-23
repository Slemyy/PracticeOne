package LinkedList

import (
	"fmt"
	"testing"
	"time"
)

func Test_AddToFront(t *testing.T) {
	blockedIPs := Create()

	testCase := []struct {
		name     string
		data     BlockedIP
		expected string
	}{
		{
			name: "Test #1",
			data: BlockedIP{
				IPAddress: "121.0.0.1",
				Reason:    "Malicious traffic",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
		{
			name: "Test #2",
			data: BlockedIP{
				IPAddress: "192.162.1.5",
				Reason:    "Brute force attack",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
		{
			name: "Test #3",
			data: BlockedIP{
				IPAddress: "10.589.548.5",
				Reason:    "for fun",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
		{
			name: "Test #4",
			data: BlockedIP{
				IPAddress: "89.109.58.8",
				Reason:    "abuse admin panel",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
		{
			name: "Test #5",
			data: BlockedIP{
				IPAddress: "192.162.1.5",
				Reason:    "Brute force attack",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
	}

	for _, i := range testCase {
		blockedIPs.AddToFront(i.data)

		t.Logf("Тест <%s> успешно отработал.", i.name)
	}

	fmt.Println()
	err := blockedIPs.Print()
	if err != nil {
		return
	}

	err = blockedIPs.DeleteFromFront()
	if err != nil {
		return
	}

	err = blockedIPs.Print()
	if err != nil {
		return
	}

}

func Test_linkedList_Search(t *testing.T) {
	blockedIPs := Create()

	//value := BlockedIP{IPAddress: "192.162.1.5", Reason: "Brute force attack"}
	value := BlockedIP{IPAddress: "89.109.58.8", Reason: "abuse admin panel"}

	testCase := []struct {
		name     string
		data     BlockedIP
		expected string
	}{
		{
			name: "Test #1",
			data: BlockedIP{
				IPAddress: "121.0.0.1",
				Reason:    "Malicious traffic",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
		{
			name: "Test #2",
			data: BlockedIP{
				IPAddress: "89.109.58.8",
				Reason:    "abuse admin panel",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
		{
			name: "Test #3",
			data: BlockedIP{
				IPAddress: "10.589.548.5",
				Reason:    "for fun",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
		{
			name: "Test #4",
			data: BlockedIP{
				IPAddress: "89.109.58.8",
				Reason:    "abuse admin panel",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
		{
			name: "Test #5",
			data: BlockedIP{
				IPAddress: "192.162.1.5",
				Reason:    "Brute force attack",
				BlockedAt: time.Now(),
			},
			expected: "",
		},
	}

	for _, i := range testCase {
		blockedIPs.AddToFront(i.data)

		t.Logf("Тест <%s> успешно отработал.", i.name)
	}

	fmt.Println()
	err := blockedIPs.Print()
	if err != nil {
		return
	}

	for i := 0; i < blockedIPs.length; i++ {
		val, err := blockedIPs.GetValue(i)
		if err != nil {
			t.Error("Error: " + err.Error())
		}

		if val.Reason == value.Reason && val.IPAddress == value.IPAddress {
			err := blockedIPs.DeleteAtIndex(i)
			if err != nil {
				t.Error("Error: " + err.Error())
			}
		}
	}

	err = blockedIPs.Print()
	if err != nil {
		return
	}
}
