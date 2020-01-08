export function CustomDateComponent () {
}

CustomDateComponent.prototype.init = function (params) {
  var template =
        "<input type='text' data-input />" +
        "<a class='input-button' title='clear' data-clear>" +
            "<i class='fa fa-times'></i>" +
        '</a>'

  this.params = params

  this.eGui = document.createElement('div')

  var eGui = this.eGui

  eGui.setAttribute('role', 'presentation')
  eGui.classList.add('ag-input-wrapper')
  eGui.classList.add('custom-date-filter')
  eGui.innerHTML = template

  this.eInput = eGui.querySelector('input')

  // eslint-disable-next-line no-undef
  this.picker = flatpickr(this.eGui, {
    onChange: this.onDateChanged.bind(this),
    dateFormat: 'd/m/Y',
    wrap: true
  })

  this.picker.calendarContainer.classList.add('ag-custom-component-popup')

  this.date = null
}

CustomDateComponent.prototype.getGui = function () {
  return this.eGui
}

CustomDateComponent.prototype.onDateChanged = function (selectedDates) {
  this.date = selectedDates[0] || null
  this.params.onDateChanged()
}

CustomDateComponent.prototype.getDate = function () {
  return this.date
}

CustomDateComponent.prototype.setDate = function (date) {
  this.picker.setDate(date)
  this.date = date
}

// TimeFilter
export function CustomTimeComponent () {
}

CustomTimeComponent.prototype.init = function (params) {
  this.valueGetter = params.valueGetter
  this.filterText = null
  this.setupGui(params)
}

// not called by ag-Grid, just for us to help setup
CustomTimeComponent.prototype.setupGui = function (params) {
  this.gui = document.createElement('div')
  this.gui.innerHTML =
        '<div style="padding: 4px; width: 200px;">' +
        '<div style="font-weight: bold;">Hour Minute</div>' +
        '<div><input style="margin: 4px 0px 4px 0px;" type="text" id="filterText" placeholder="HH:mm"/></div>' +
        '</div>'

  this.eFilterText = this.gui.querySelector('#filterText')
  this.eFilterText.addEventListener('changed', listener)
  this.eFilterText.addEventListener('paste', listener)
  this.eFilterText.addEventListener('input', listener)
  // IE doesn't fire changed for special keys (eg delete, backspace), so need to
  // listen for this further ones
  this.eFilterText.addEventListener('keydown', listener)
  this.eFilterText.addEventListener('keyup', listener)

  var that = this
  function listener (event) {
    that.filterText = event.target.value
    params.filterChangedCallback()
  }
}

CustomTimeComponent.prototype.getGui = function () {
  return this.gui
}

CustomTimeComponent.prototype.doesFilterPass = function (params) {
  // make sure each word passes separately, ie search for firstname, lastname
  var passed = true
  var valueGetter = this.valueGetter
  this.filterText.toLowerCase().split(' ').forEach(function (filterWord) {
    var value = valueGetter(params)
    if (value.toString().toLowerCase().indexOf(filterWord) < 0) {
      passed = false
    }
  })

  return passed
}

CustomTimeComponent.prototype.isFilterActive = function () {
  return this.filterText !== null && this.filterText !== undefined && this.filterText !== ''
}

CustomTimeComponent.prototype.getModel = function () {
  var model = { value: this.filterText.value }
  return model
}

CustomTimeComponent.prototype.setModel = function (model) {
  this.eFilterText.value = model.value
}
