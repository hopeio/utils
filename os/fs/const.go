/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package fs

import "os"

const PathSeparator = string(os.PathSeparator)

type FileType int

const (
	Unknown FileType = iota
	Txt
	Doc
	Docx
	Xls
	Xlsx
)

var FileTypeMap = map[string]FileType{
	".txt": Txt,
	".doc": Doc,
}
