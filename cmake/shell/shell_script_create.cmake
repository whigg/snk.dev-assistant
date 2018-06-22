
# creates a shell script file containing the specified code and the correct extesion to execute
# with execute_process
function(shell_script_create path code)
  if(NOT ARGN)
    shell_get()
    ans(shell)
  else()
    set(shell "${ARGN}")
  endif()
  if("${shell}_" STREQUAL "cmd_")
    if(NOT "${path}" MATCHES "\\.bat$")
      set(path "${path}.bat")
    endif()
    set(code "@echo off\n${code}")
  elseif("${shell}_" STREQUAL "bash_")
    if(NOT "${path}" MATCHES "\\.sh$")
      set(path "${path}.sh")
    endif()
    set(code "#!/bin/bash\n${code}")
    touch("${path}")
    execute_process(COMMAND chmod +x "${path}")
  else()
    message(WARNING "shell not supported: '${shell}' ")
    return()
  endif()
    fwrite("${path}" "${code}")
    return_ref(path)
endfunction()