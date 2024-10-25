package godopedb;

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Get value from the db. Returns nil as ValueResponse if the provided value is not present in the namespace. Returns error if there is an error while fetching the data or provided namespace does not exist
func GetValue(db_url string, options SelectValueRequest) (*ValueResponse, error) {
	url := fmt.Sprintf("%s/api/get", db_url);
	body, err := json.Marshal(options);
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body));
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json");
 	client := &http.Client{}
    	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close();
	if resp.StatusCode == 404 {
		return nil, nil
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid namespace options")
	} else {
		res := &ValueResponse{};
		body, err := io.ReadAll(resp.Body);
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, res);
		return res, err
	}
}

// Insert specifieg KV pair into a namespace. Returns new DB state checksum after value insertion, or error indicating communication failure or that the requested namespace does not exist
func InsertValue(db_url string, options InsertValueRequest) (*ChecksumResponse, error) {
	url := fmt.Sprintf("%s/api/insert", db_url);
	body, err := json.Marshal(options);
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body));
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json");
 	client := &http.Client{}
    	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close();
	if resp.StatusCode == 404 {
		return nil, nil
	} else if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Invalid namespace options")
	} else {
		res := &ChecksumResponse{};
		body, err := io.ReadAll(resp.Body);
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, res);
		return res, err
	}
}


// Deletes the KV pair associated with the provided key in provided namespace. Returns the checksum of the entire db state after key deletion (does not delete namespace, even if empty);
func DeleteValue(db_url string, options SelectValueRequest) (*ChecksumResponse, error) {
	url := fmt.Sprintf("%s/api/delete", db_url);
	body, err := json.Marshal(options);
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(body));
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json");
 	client := &http.Client{}
    	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close();
	if resp.StatusCode == 404 {
		return nil, nil
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid namespace options")
	} else {
		res := &ChecksumResponse{};
		body, err := io.ReadAll(resp.Body);
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, res);
		return res, err
	}
}

// Obtains the checksum of the entire db state. Returns error if db is uninitialized or a communication error occurs
func GetChecksum(db_url string) (*ChecksumResponse, error) {
	url := fmt.Sprintf("%s/api/checksum", db_url);
	req, err := http.NewRequest("GET", url, nil);
	if err != nil {
		return nil, err
	}
 	client := &http.Client{}
    	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close();
	if resp.StatusCode != 200 {
		return nil, nil
	} else {
		res := &ChecksumResponse{};
		body, err := io.ReadAll(resp.Body);
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, res);
		return res, err
	}
}

// Creates a new namespace in the db. This operation does not update the checksum. This operation is idempotent. Creating a namespace automatically creates any parent namespaces if absent.
func CreateNamespace(db_url string, options SelectNamespaceRequest) error {
	url := fmt.Sprintf("%s/api/namespace", db_url);
	body, err := json.Marshal(options);
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body));
	if err != nil {
		return err
	}
	req.Header.Set("content-type", "application/json");
 	client := &http.Client{}
    	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close();
	if resp.StatusCode != 201 {
		return fmt.Errorf("Could not create namespace")
	} else {
		return nil
	}
}
