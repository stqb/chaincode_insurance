/**
 * Created by Ditty on 16/5/23.
 */

var http = require('http');

var DEV = true;

function generateChainCodeName(chaincode, id) {
    return "BC-NXY_" + chaincode + "_" + id;
}

function generateChainCodePath(chaincode) {
    return "/opt/gopath/src/github.com/hyperledger/fabric/BC-NXY/chaincode/" + chaincode;// + "/" + chaincode;
}

function prepareInit(chaincode, id) {
    var cmd_sign = "CORE_CHAINCODE_ID_NAME=" + generateChainCodeName(chaincode, id) + " CORE_PEER_ADDRESS=0.0.0.0:30303 /opt/gopath/src/github.com/hyperledger/fabric/BC-NXY/chaincode/sign/sign";
    var cprocess = require('child_process');
    cprocess.exec(cmd_sign, function (err, stdout, stderr){
        console.log("stdout: " + stdout);
        console.log("stderr: " + stderr);
        if (err !== null) {
            console.log("exec error: " + err);
        }
    });
}

function config(hostname, portNum, pathURL) {
    var host = "localhost";
    if (hostname) {
        host = hostname;
    }

    var port = 5000;
    if (portNum) {
        port = portNum;
    }

    var path = '/chaincode';
    if (pathURL) {
        path = pathURL;
    }

    return options = {
        hostname: host,
        port: port,
        path: path,
        method: 'POST',
        headers: {
            'Content-Type': 'text/html'
        }
    };
}

function prepare(chaincode, id, func, args) {
    var post_data = {
        "jsonrpc": "2.0",
        "params": {
            "type": 1,
            "chaincodeID":{
                "path": generateChainCodePath(chaincode, id)
            },
            "ctorMsg": {
                "function":func,
                "args":args
            }
        }
    };

    if (DEV) {
        post_data = {
            "jsonrpc": "2.0",
            "params": {
                "type": 1,
                "chaincodeID":{
                    "name": generateChainCodeName(chaincode, id)
                },
                "ctorMsg": {
                    "function":func,
                    "args":args
                }
            }
        };
    }

    if(func === 'init') {
        post_data["method"] = "deploy";
        post_data["id"] = 1;
    } else {
        post_data["method"] = func;
        post_data["id"] = 3;
    }

    return JSON.stringify(post_data);
}

function call(chaincode, id, func, args, nextFunc, res) {
    if(func === 'init') {
        prepareInit(chaincode, id);
    }

    var postStr = prepare(chaincode, id, func, args);
    console.log('postStr:' + postStr);

    var options = config();
    options['headers']['Content-Length'] = postStr.length;
    console.log('options:' + options);

    var req = http.request(options, function(response) {
        console.log('STATUS: ' + response.statusCode);
        console.log('HEADERS: ' + JSON.stringify(response.headers));
        response.setEncoding('utf8');
        var resdata = '';
        response.on('data', function (chunk) {
            resdata += chunk;
        });
        response.on('end', function() {
            console.log('DATA: ' + resdata);
            if(response.statusCode == 200) {
                var jsonData = JSON.parse(resdata);
                if(jsonData.result){
                    nextFunc(resdata);
                    return;
                }
            }
            nextFunc(null, {statusCode: response.statusCode, statusMessage: response.statusMessage});
        });
    });

    req.on('error', function(err) {
        console.log('problem with request: ' + err.message);
        // res.writeHead(500, {'Content-Type': 'text/html', 'Access-Control-Allow-Origin': originURL, 'Access-Control-Request-Method': 'GET, POST, PUT'});
        // res.end(err.message);
        // return;
        nextFunc(null, err);
    });

    req.write(postStr);
    console.log('req.write(postStr) is OK');
    req.end();
}

exports.call = call;