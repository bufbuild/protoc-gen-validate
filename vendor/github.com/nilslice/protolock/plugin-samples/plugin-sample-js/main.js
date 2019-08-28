// your plugin receives JSON into its stdin, in the shape of the data object below
let data = {
    current: {},
    updated: {},
    protolock_warnings: [{
        filepath: "",
        message: "",
        name: "",
    }],
    plugin_warnings: [{
        filepath: "",
        message: "",
        name: "",
    }],
    plugin_error_message: "",
}

process.stdin.setEncoding('utf8')

function readStdinSync() {
    return new Promise((resolve, reject) => {
        process.stdin.resume()
        process.stdin.on('data', function (data) {
            process.stdin.pause()
            resolve(data)
        })
    })
}

async function main() {
    data = JSON.parse(await readStdinSync())

    console.log(JSON.stringify(customRuleFunc(data))) // console.log writes to stdout
}

function customRuleFunc(data) {
    // compare proto primitives inside the data.current and data.updated objects
    let warnings = [
        {
            filepath: "path/to/file.proto",
            message: "Something bad happened."
        },
        {
            filepath: "path/to/another.proto",
            message: "Something else bad happened."
        }

    ]

    data.plugin_warnings = (data.plugin_warnings || []).concat(warnings)
    return data
}

main()