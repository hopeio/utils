/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package param

type IPageSort interface {
	IPage
	ISort
}

type IPage interface {
	PageNo() int
	PageSize() int
}

type ISort interface {
	Column() string
	Type() SortType
}
