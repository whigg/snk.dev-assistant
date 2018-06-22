# returns the time elapsed since timer identified by id was started
function(timer_elapsed id)      
  millis()
  ans(now)
  map_get(__timers ${id})
  ans(then)
  # so this has to be done because cmake can't handle numbers which are too large
  string_trim_to_difference(then now)
  map_tryget(__timers __prejudice)      
  ans(prejudice)
  ## if now and the are equal null is returned so here it normalized
  if(NOT now)
    set(now 0)
  endif()
  if(NOT then)
    set(then 0)
  endif()
  math(EXPR elapsed "${now} - ${then} - ${prejudice}")
  math_max(${elapsed} 0)
  ans(elapsed)
  return("${elapsed}")
endfunction()