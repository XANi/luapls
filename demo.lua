local num = 3
local other_num = num
local other_other_num = other_num
local other_other_other_num = other_other_num

--- @type number
local uninitialized

local str, bool = "foo", true
local unknown = nil
local func = function(foo, bar) end

local type_error = num + str -- Cannot add a 'number' with a 'string'.
