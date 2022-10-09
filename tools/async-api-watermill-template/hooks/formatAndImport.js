const {exec} = require("child_process");

module.exports = {
    'generate:after': formatAndImport
};

async function formatAndImport(generator) {
    // execute gofmt and goimports for all go files in target dir
    exec("find " + generator.targetDir + " -name 'asyncapi_*.go' | while read -r file; do gofmt -w -s \"$file\"; goimports -w \"$file\"; done", (error, stdout, stderr) => {
        if (error) {
            console.log(`error: ${error.message}`);
        }
        if (stderr) {
            console.log(`stderr: ${stderr}`);
        }
    })
}