package keycodes

var Code map[string]int
//All keys that are here are taken from here: https://github.com/LuaDist/sdl/blob/master/include/SDL_scancode.h
func init(){
  Code = make(map[string]int)
  Code["a"] = 4
  Code["b"] = 5
  Code["c"] = 6
  Code["d"] = 7
  Code["e"] = 8
  Code["f"] = 9
  Code["g"] = 10
  Code["h"] = 11
  Code["i"] = 12
  Code["j"] = 13
  Code["k"] = 14
  Code["l"] = 15
  Code["m"] = 16
  Code["n"] = 17
  Code["o"] = 18
  Code["p"] = 19
  Code["q"] = 20
  Code["r"] = 21
  Code["s"] = 22
  Code["t"] = 23
  Code["u"] = 24
  Code["v"] = 25
  Code["w"] = 26
  Code["x"] = 27
  Code["y"] = 28
  Code["z"] = 29

  Code["1"] = 30
  Code["2"] = 31
  Code["3"] = 32
  Code["4"] = 33
  Code["5"] = 34
  Code["6"] = 35
  Code["7"] = 36
  Code["8"] = 37
  Code["9"] = 38
  Code["0"] = 39

  Code["enter"] = 40
  //Code["return"] = 40
  Code["escape"] = 41
  Code["backspace"] = 42
  Code["tab"] = 43
  Code["space"] = 44

  Code["minus"] = 45
  Code["equals"] = 46
  Code["left_bracket"] = 47
  Code["right_bracket"] = 48
  Code["backlash"] = 49

  Code["nonushash"] = 50

  Code["semicolon"] = 51
  Code["apostrophe"] = 52
  Code["grave"] = 53

  Code["comma"] = 54
  Code["period"] = 55
  Code["slash"] = 56

  Code["capslock"] = 57

  Code["f1"] = 58
  Code["f2"] = 59
  Code["f3"] = 60
  Code["f4"] = 61
  Code["f5"] = 62
  Code["f6"] = 63
  Code["f7"] = 64
  Code["f8"] = 65
  Code["f9"] = 66
  Code["f10"] = 67
  Code["f11"] = 68
  Code["f12"] = 69

  Code["printscreen"] = 70
  Code["scrolllock"] = 71
  Code["pause"] = 72
  Code["insert"] = 73
  Code["home"] = 74
  Code["pageup"] = 75
  Code["delete"] = 76
  Code["end"] = 77
  Code["pagedown"] = 78
  Code["right"] = 79
  Code["left"] = 80
  Code["down"] = 81
  Code["up"] = 82

  Code["numlockclear"] = 83

  Code["kp_divide"] = 84
  Code["kp_multiply"] = 85
  Code["kp_minus"] = 86
  Code["kp_plus"] = 87
  Code["kp_enter"] = 88
  Code["kp_1"] = 89
  Code["kp_2"] = 90
  Code["kp_3"] = 91
  Code["kp_4"] = 92
  Code["kp_5"] = 93
  Code["kp_6"] = 94
  Code["kp_7"] = 95
  Code["kp_8"] = 96
  Code["kp_9"] = 97
  Code["kp_0"] = 98
  Code["kp_period"] = 99

  Code["nonus_backslash"] = 100

  Code["application"] = 101 /**< windows contextual menu compose */
  Code["power"] = 102 /**< The USB document says this is a status flag
                             *   not a physical key - but some Mac keyboards
                             *   do have a power key. */
  Code["kp_equals"] = 103
  Code["f13"] = 104
  Code["f14"] = 105
  Code["f15"] = 106
  Code["f16"] = 107
  Code["f17"] = 108
  Code["f18"] = 109
  Code["f19"] = 110
  Code["f20"] = 111
  Code["f21"] = 112
  Code["f22"] = 113
  Code["f23"] = 114
  Code["24"] = 115
  Code["execute"] = 116
  Code["help"] = 117
  Code["menu"] = 118
  Code["select"] = 119
  Code["stop"] = 120
  Code["again"] = 121   /**< redo */
  Code["undo"] = 122
  Code["cut"] = 123
  Code["copy"] = 124
  Code["paste"] = 125
  Code["find"] = 126
  Code["mute"] = 127
  Code["volume_up"] = 128
  Code["volume_down"] = 129
/* not sure whether there's a reason to enable these */
/*     SDL_SCANCODE_LOCKINGCAPSLOCK = 130  */
/*     SDL_SCANCODE_LOCKINGNUMLOCK = 131 */
/*     SDL_SCANCODE_LOCKINGSCROLLLOCK = 132 */
  Code["kp_comma"] = 133
  Code["kp_equals_400"] = 134

  Code["international1"] = 135 /**< used on Asian keyboards see
                                          footnotes in USB doc */
  Code["international2"] = 136
  Code["international3"] = 137 /**< Yen */
  Code["international4"] = 138
  Code["international5"] = 139
  Code["international6"] = 140
  Code["international7"] = 141
  Code["international8"] = 142
  Code["international9"] = 143
  Code["lang1"] = 144 /**< Hangul/English toggle */
  Code["lang2"] = 145 /**< Hanja conversion */
  Code["lang3"] = 146 /**< Katakana */
  Code["lang4"] = 147 /**< Hiragana */
  Code["lang5"] = 148 /**< Zenkaku/Hankaku */
  Code["lang6"] = 149 /**< reserved */
  Code["lang7"] = 150 /**< reserved */
  Code["lang8"] = 151 /**< reserved */
  Code["lang9"] = 152 /**< reserved */

  Code["alterase"] = 153 /**< Erase-Eaze */
  Code["sysreq"] = 154
  Code["cancel"] = 155
  Code["clear"] = 156
  Code["prior"] = 157
  Code["enter2"] = 158
  Code["separator"] = 159
  Code["out"] = 160
  Code["oper"] = 161
  Code["clearagain"] = 162
  Code["crsel"] = 163
  Code["exsel"] = 164

  Code["kp_00"] = 176
  Code["kp_000"] = 177
  Code["thousand_separator"] = 178
  Code["decimal_separator"] = 179
  Code["currency_unit"] = 180
  Code["currency_b_unit"] = 181
  Code["kp_left_paren"] = 182
  Code["kp_right_paren"] = 183
  Code["kp_left_brace"] = 184
  Code["kp_right_brace"] = 185
  Code["kp_tab"] = 186
  Code["kp_backspace"] = 187
  Code["kp_a"] = 188
  Code["kp_b"] = 189
  Code["kp_c"] = 190
  Code["kp_d"] = 191
  Code["kp_e"] = 192
  Code["kp_f"] = 193
  Code["kp_xor"] = 194
  Code["kp_power"] = 195
  Code["kp_percent"] = 196
  Code["kp_less"] = 197
  Code["kp_greater"] = 198
  Code["kp_ampersand"] = 199
  Code["kp_dbl_ampersand"] = 200
  Code["kp_vertical_bar"] = 201
  Code["kp_dbl_vertical_bar"] = 202
  Code["kp_colon"] = 203
  Code["kp_hash"] = 204
  Code["kp_space"] = 205
  Code["kp_at"] = 206
  Code["kp_exclam"] = 207
  Code["kp_mem_store"] = 208
  Code["kp_mem_recall"] = 209
  Code["kp_mem_clear"] = 210
  Code["kp_mem_add"] = 211
  Code["kp_mem_sunstract"] = 212
  Code["kp_mem_multiply"] = 213
  Code["kp_mem_divide"] = 214
  Code["kp_plusminus"] = 215
  Code["kp_clear"] = 216
  Code["kp_clear_entry"] = 217
  Code["kp_binary"] = 218
  Code["kp_octal"] = 219
  Code["kp_decimal"] = 220
  Code["kp_hexdecimal"] = 221

  Code["lctrl"] = 224
  Code["lshift"] = 225
  Code["lalt"] = 226 /**< alt option */
  Code["lgui"] = 227 /**< windows command (apple) meta */
  Code["rctrl"] = 228
  Code["rshift"] = 229
  Code["ralt"] = 230 /**< alt gr option */
  Code["rgui"] = 231 /**< windows command (apple) meta */

  Code["mode"] = 257    /**< I'm not sure if this is really not covered
                               *   by any of the above but since there's a
                               *   special KMOD_MODE for it I'm adding it here
                               */

  /* @} *//* Usage page 0x07 */

  /**
   *  \name Usage page 0x0C
   *
   *  These values are mapped from usage page 0x0C (USB consumer page).
   */
  /* @{ */

  Code["audio_next"] = 258
  Code["audio_prev"] = 259
  Code["audio_stop"] = 260
  Code["audio_play"] = 261
  Code["audio_mute"] = 262
  Code["media_select"] = 263
  Code["www"] = 264
  Code["mail"] = 265
  Code["calculator"] = 266
  Code["computer"] = 267
  Code["ac_search"] = 268
  Code["ac_home"] = 269
  Code["ac_back"] = 270
  Code["ac_forward"] = 271
  Code["ac_stop"] = 272
  Code["ac_refresh"] = 273
  Code["ac_bookmarks"] = 274

  /* @} *//* Usage page 0x0C */

  /**
   *  \name Walther keys
   *
   *  These are values that Christian Walther added (for mac keyboard?).
   */
  /* @{ */

  Code["brightness_down"] = 275
  Code["brightness_up"] = 276
  Code["display_switch"] = 277 /**< display mirroring/dual display
                                         switch video mode switch */
  Code["kbd_illum_toggle"] = 278
  Code["kbd_illum_down"] = 279
  Code["eject"] = 281
  Code["sleep"] = 282

  Code["app1"] = 283
  Code["app2"] = 284
}
