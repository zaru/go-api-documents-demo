const hooks = require('hooks')
const execSync = require('child_process').execSync

const testDBHost = 'db'
const testDBName = 'sample_test'
const testAppPort = '1323'

hooks.beforeAll(function (transactions, done) {
  hooks.log('before all')

  let result = execSync('mysql -h ' + testDBHost + ' -u root -e "create database ' + testDBName + '"').toString()
  hooks.log('create database: ' + result)

  result = execSync('goose -dir ./migrations mysql "root@(' + testDBHost + ')/' + testDBName + '" up').toString()
  hooks.log('migration: ' + result)

  result = execSync('mysql -h ' + testDBHost + ' -u root ' + testDBName + ' < ./seed/test.sql').toString()
  hooks.log('seed: ' + result)

  done()
})

hooks.beforeEach(function (transaction, done) {
  hooks.log('before each')
  done()
})

hooks.before("Machines > Machines collection > Get Machines", function (transaction, done) {
  hooks.log("before")
  done()
})

hooks.beforeEachValidation(function (transaction, done) {
  hooks.log('before each validation')
  done()
})

hooks.beforeValidation("Machines > Machines collection > Get Machines", function (transaction, done) {
  hooks.log("before validation")
  done()
})

hooks.after("Machines > Machines collection > Get Machines", function (transaction, done) {
  hooks.log("after")
  done()
})

hooks.afterEach(function (transaction, done) {
  hooks.log('after each')
  done()
})

hooks.afterAll(function (transactions, done) {
  hooks.log('after all')

  let result = execSync('mysql -h ' + testDBHost + ' -u root -e "drop database ' + testDBName + '"').toString()
  hooks.log('drop database: ' + result)

  result = execSync('kill `lsof -ti tcp:' + testAppPort + '`')
  hooks.log('kill app: ' + result)

  done()
})
