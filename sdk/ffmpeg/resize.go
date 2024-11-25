/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package ffmpeg

const ResizeCmd = CommonCmd + `-vf "scale=iw*.5:ih*.5" %s`
