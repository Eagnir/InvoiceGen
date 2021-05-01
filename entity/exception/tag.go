package exception

import "errors"

var Tag_RecordNotFound = errors.New("Tag not found")

var Tag_RequiredField_Name = errors.New("Please provide a name for this tag")

var Tag_PrimeryKeyNotBlank = errors.New("Tag's primary key (name) is not blank")

var Tag_NameAlreadyExist = errors.New("Tag name already exist")
