
function(project_unloader project)
  map_tryget(${project} dependency_descriptor)
  ans(dependency_descriptor)
  map_tryget("${dependency_descriptor}" is_ready)
  ans(is_ready)
  if(NOT is_ready)
    return()
  endif()
  if("${ARGN}_" STREQUAL "_" OR "${project}" STREQUAL "${ARGN}")
    project_unload(${project})
  endif()
endfunction()