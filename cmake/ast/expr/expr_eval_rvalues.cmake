function(expr_eval_expression)
#  message("evaluating expression")
  map_get(${ast}  children)
  ans(children)
    map_new()
    ans(new_context)
  map_set(${new_context} parent_context ${context})
  map_tryget(${context}  scope)
  ans(scope)
  map_set(${new_context} scope ${scope})

  foreach(rvalue_ast ${children})
    ast_eval(${rvalue_ast} ${new_context})
    ans(rvalue)
    map_set(${new_context} left ${rvalue})
    map_set(${new_context} left_ast ${rvalue_ast})
  endforeach()


  map_tryget(${new_context}  left)
  ans(left)
  return_ref(left)
endfunction()