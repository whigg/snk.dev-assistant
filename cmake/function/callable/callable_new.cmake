function(callable_new input)

  map_new()
  ans(callable)
  map_set_special("${callable}" "$type" "callable")
  function_import("${input}")
  ans(callable_func)
  map_set_special("${callable}" callable_function "${callable_func}")
  map_set_special("${callable}" callable_input "${input}" )
  return_ref(callable)
endfunction()