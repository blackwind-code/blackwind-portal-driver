package openstack

import (
	"math/rand"
	"os/exec"

	"github.com/alessio/shellescape"
)

type HttpRes struct {
	status string
	body   string
}

func randomString() string {
	letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str := make([]rune, 16)
	for i := range str {
		str[i] = letters[rand.Intn(len(str))]
	}
	return string(str)
}

func UserCreate(email string, passwordHash string) HttpRes {
	// check username against PROHIBITED
	for _, username := range PROHIBITED {
		if email == username {
			return HttpRes{status: "403 Forbidden", body: "This username is reserved for special purpose."}
		}
	}

	argEmail := shellescape.Quote(email)
	argRandomPass := shellescape.Quote(randomString())

	// openstack project create --domain default --description <email> <email>
	// output, err := exec.Command("openstack project create --domain default --description user " + argEmail).Output()
	output, err := exec.Command("openstack", "project", "create", "--domain", "default", "--description", "user", argEmail).Output()
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return HttpRes{status: "500 Internal Server Error", body: string(output)}
	}

	// openstack user create --project <email> --password <randompass> <email>
	// output, err = exec.Command("openstack user create --project " + argEmail + " --password " + argRandomPass + " " + argEmail + " --enable-lock-password --email " + argEmail).Output()
	output, err = exec.Command("openstack", "user", "create", "--domain", "default", "--project", argEmail, "--project-domain", "default", "--password", argRandomPass, argEmail, " --enable-lock-password", "--email", argEmail).Output()
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return HttpRes{status: "500 Internal Server Error", body: string(output)}
	}

	output, err = exec.Command("openstack", "role", "add", "--project", argEmail, "--user", argEmail, "--user-domain", "default", "member").Output()
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return HttpRes{status: "500 Internal Server Error", body: string(output)}
	}

	// UPDATE password SET password_hash='<password_hash>' WHERE local_user_id=(SELECT id FROM local_user WHERE name='<email>')
	_, err = stmtUpdateUserPasswordHash.Exec(passwordHash, email)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return HttpRes{status: "500 Internal Server Error", body: err.Error()}
	}

	return HttpRes{status: "200 OK", body: ""}
}

func UserUpdate(oldEmail string, newEmail string, passwordHash string) HttpRes {
	// check username against PROHIBITED
	for _, username := range PROHIBITED {
		if newEmail == username {
			return HttpRes{status: "403 Forbidden", body: "This username is reserved for special purpose."}
		}
	}

	argOldEmail := shellescape.Quote(oldEmail)
	argNewEmail := shellescape.Quote(newEmail)

	// change project name
	//output, err := exec.Command("openstack project set " + argOldEmail + " --name " + argNewEmail).Output()
	output, err := exec.Command("openstack", "project", "set", argOldEmail, "--name", argNewEmail).Output()
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return HttpRes{status: "500 Internal Server Error", body: string(output)}
	}

	// change user email
	// output, err = exec.Command("openstack user set " + argOldEmail + " --name " + argNewEmail).Output()
	output, err = exec.Command("openstack", "user", "set", argOldEmail, "--name", argNewEmail).Output()
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return HttpRes{status: "500 Internal Server Error", body: string(output)}
	}

	// change user password
	_, err = stmtUpdateUserPasswordHash.Exec(passwordHash, newEmail)
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return HttpRes{status: "500 Internal Server Error", body: err.Error()}
	}

	return HttpRes{status: "200 OK", body: ""}
}

func UserDelete(email string) HttpRes {
	// check username against PROHIBITED
	for _, username := range PROHIBITED {
		if email == username {
			return HttpRes{status: "403 Forbidden", body: "This username is reserved for special purpose."}
		}
	}

	argEmail := shellescape.Quote(email)

	// delete user
	// output, err := exec.Command("openstack user delete " + argEmail).Output()
	output, err := exec.Command("openstack", "user", "delete", argEmail).Output()
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return HttpRes{status: "500 Internal Server Error", body: string(output)}
	}

	// delete project
	// output, err = exec.Command("openstack project delete " + argEmail).Output()
	output, err = exec.Command("openstack", "project", "delete", argEmail).Output()
	if err != nil {
		Log.Printf("Error: %v\n", err.Error())
		return HttpRes{status: "500 Internal Server Error", body: string(output)}
	}

	return HttpRes{status: "200 OK", body: ""}
}
