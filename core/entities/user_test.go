package entities

import (
    "testing"
)

func TestFindUser(t *testing.T) {
    _, createErr := CreateUser(&User{Token: "mytoken"})
    if (createErr != nil) {
        t.Error(createErr)
    }

    user, findErr := FindUserByToken("mytoken")
    if (findErr != nil) {
        t.Error(findErr)
    }
    if (user.Token != "mytoken"){
        t.Error("it should find a user with token 'mytoken'")
    }

    removeErr := RemoveUser(user)
    if (removeErr != nil) {
        t.Error(removeErr)
    }
}
