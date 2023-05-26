// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"

/**
 * This function uses a selection component from https://github.com/CreativeIT/getmdl-select
 */

function myButtonClicked() {
  // input
  const age = parseInt(document.getElementById('age-entered').value)
  const day = document.getElementById('day-selected').value

  // process
  if (day == "Tuesday" || day == "Thursday" || age > 12 && age < 21) {
  document.getElementById('answer').innerHTML = "You get the student discount."
  } else {
    document.getElementById('answer').innerHTML = "You must pay regular price."
  }
}
