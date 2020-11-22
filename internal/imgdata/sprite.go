package imgdata

import "../img"

var XshipSprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 26,
				Y: 0,
			},
		},
		Width:  26,
		Height: 16,
	},
	HitBox: img.HitBox{
		Offset: img.Position{
			X: 1,
			Y: 3,
		},
		Width:  23,
		Height: 12,
	},
}

var MissileSprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{{
			X: 0,
			Y: 16,
		}},
		Width:  10,
		Height: 3,
	},
	HitBox: img.HitBox{
		Offset: img.Position{
			X: 1,
			Y: 0,
		},
		Width:  8,
		Height: 3,
	},
}

var BadGuy1Sprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 0,
				Y: 19,
			},
			{
				X: 21,
				Y: 19,
			},
			{
				X: 42,
				Y: 19,
			},
		},
		Width:  21,
		Height: 16,
	},
	HitBox: img.HitBox{
		Offset: img.Position{
			X: 4,
			Y: 2,
		},
		Width:  16,
		Height: 12,
	},
}

var ExplodedBadGuy1Sprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 0,
				Y: 35,
			},
			{
				X: 21,
				Y: 35,
			},
			{
				X: 42,
				Y: 35,
			},
		},
		Width:  21,
		Height: 16,
	},
}

var BadGuy2Sprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 0,
				Y: 92,
			},
			{
				X: 15,
				Y: 92,
			},
			{
				X: 30,
				Y: 92,
			},
		},
		Width:  15,
		Height: 12,
	},
	HitBox: img.HitBox{
		Offset: img.Position{
			X: 2,
			Y: 1,
		},
		Width:  12,
		Height: 10,
	},
}

var ExplodedBadGuy2Sprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 0,
				Y: 104,
			},
			{
				X: 15,
				Y: 104,
			},
			{
				X: 30,
				Y: 104,
			},
		},
		Width:  15,
		Height: 12,
	},
}

var BadMissileSprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 34,
				Y: 74,
			},
			{
				X: 38,
				Y: 74,
			},
			{
				X: 42,
				Y: 74,
			},
		},
		Width:  4,
		Height: 4,
	},
	HitBox: img.HitBox{
		Offset: img.Position{
			X: 1,
			Y: 1,
		},
		Width:  2,
		Height: 2,
	},
}

var PressStartSprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 0,
				Y: 51,
			},
		},
		Width:  45,
		Height: 9,
	},
}

var SkullSprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 0,
				Y: 60,
			},
		},
		Width:  34,
		Height: 32,
	},
}

var NumberSprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 34,
				Y: 60,
			},
			{
				X: 37,
				Y: 60,
			},
			{
				X: 40,
				Y: 60,
			},
			{
				X: 43,
				Y: 60,
			},
			{
				X: 46,
				Y: 60,
			},
			{
				X: 49,
				Y: 60,
			},
			{
				X: 52,
				Y: 60,
			},
			{
				X: 55,
				Y: 60,
			},
			{
				X: 58,
				Y: 60,
			},
			{
				X: 61,
				Y: 60,
			},
		},
		Width:  3,
		Height: 5,
	},
}

var GlowingWhiteStarSprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 34,
				Y: 65,
			},
			{
				X: 43,
				Y: 65,
			},
			{
				X: 52,
				Y: 65,
			},
			{
				X: 61,
				Y: 65,
			},
			{
				X: 61,
				Y: 65,
			},
			{
				X: 52,
				Y: 65,
			},
			{
				X: 43,
				Y: 65,
			},
			{
				X: 34,
				Y: 65,
			},
		},
		Width:  9,
		Height: 9,
	},
}

var PauseSprite = img.Sprite{
	ImageAsset: img.ImageAsset{
		Asset: &asset1,
		Offset: []img.Position{
			{
				X: 45,
				Y: 51,
			},
		},
		Width:  23,
		Height: 9,
	},
}
