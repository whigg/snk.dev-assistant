
  function(ast_parse_regex definition stream create_node)
    nav(regex = "definition.regex")
    # regex - try match
    if(NOT regex)
      return(false)
    endif()
   # message("regex: ${regex}")
    stream_take_regex(${stream} "${regex}")
    ans(match)
  #  message("matched: '${match}'")
    if(NOT match)
      return(false)
    endif()
    nav(replace = definition.replace)
    if(replace)
      #message("replace: ${replace}")
      string(REGEX REPLACE "${regex}" "\\${replace}" match "${match}")
    endif()
    if(NOT create_node)
     # message("create_node: ${create_node}")
      return(true)
    endif()
    map_new()
    ans(node)
    map_set(${node} data "${match}")
    return(${node})
  endfunction()