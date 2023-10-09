package compiled

import "strconv"

func ProtoUserToValueMyDB(user *User) map[string]*Value {

	data := make(map[string]*Value)

	data["username"] = &Value{
		Value:  user.GetUsername(),
		Unique: true,
	}
	data["password"] = &Value{
		Value:  user.GetPassword(),
		Unique: false,
	}
	data["email"] = &Value{
		Value:  user.GetEmail(),
		Unique: false,
	}
	data["admin"] = &Value{
		Value:  strconv.FormatBool(user.GetAdmin()),
		Unique: false,
	}

	return data
}

func MyDBValueToProtoUser(value map[string]*Value) *User {
	var user *User

	boolValue, _ := strconv.ParseBool(value["admin"].Value)

	user.Id = &value["id"].Value
	user.Username = &value["username"].Value
	user.Email = &value["email"].Value
	user.Password = &value["password"].Value
	user.Admin = &boolValue

	return user
}

func MyDBValueToProtoUserList(allRows *AllRows) *UserList {
	var userList *UserList

	for _, v := range allRows.Rows {
		var u *User
		boolValue, _ := strconv.ParseBool(v.Data["admin"].Value)

		u.Id = &v.Data["id"].Value
		u.Username = &v.Data["username"].Value
		u.Email = &v.Data["email"].Value
		u.Password = &v.Data["password"].Value
		u.Admin = &boolValue

		userList.Users = append(userList.Users, u)
	}

	return userList
}
