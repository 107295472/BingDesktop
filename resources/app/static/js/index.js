let index = {
    about: function(html) {
        let c = document.createElement("div");
        c.innerHTML = html;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    addFolder(name, path) {
        let div = document.createElement("div");
        div.className = "dir";
        div.onclick = function() { index.explore(path) };
        div.innerHTML = `<i class="fa fa-folder"></i><span>` + name + `</span>`;
        document.getElementById("dirs").appendChild(div)
    },
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();

            // Explore default path
            index.Test(true);
        })
    },
    SaveImg:function (filename) {
        //filename=document.getElementById("saveImg").getAttribute("picurl");
        astilectron.showSaveDialog({title: "保存为"}, function(filename) {
            console.log("chosen filename is ",filename);
        })
    },
    Test:function (isFirst) {
        // Create message
        let message = {"name": "bing"};
        if (typeof path !== "undefined") {
            message.payload = path
        }
        var idx=document.getElementById("idx").value;
        //console.log(idx);
        if(isFirst==false)
        {
            idx++;
           // console.log(idx+"加1");
        }

        message.payload =idx+"";
        if(idx<10)
        {
            // Send message
            asticode.loader.show();
            astilectron.sendMessage(message, function(message) {
                // Init
                // Check error
                if (message.name === "error") {
                    asticode.notifier.error(message.payload);
                    return
                }
                document.getElementById("idx").value=message.payload.idx;
                document.getElementById("ti").innerText=message.payload.ti;
                document.getElementById("bingimg").src=message.payload.img;
                document.getElementById("saveImg").setAttribute("picurl",message.payload.img);
            })
            var timer = setInterval(function(){
                if (document.getElementById('bingimg').complete){
                    clearInterval(timer);
                    asticode.loader.hide();
                }
            }, 10);
        }else
            asticode.notifier.info("没有图片了");
    },
    explore: function(path) {
        // Create message
        let message = {"name": "explore"};
        if (typeof path !== "undefined") {
            message.payload = path
        }

        // Send message
        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            // Init
            asticode.loader.hide();

            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }

            // // Process path
            //document.getElementById("path").innerHTML = message.payload.path;
            //
            // // Process dirs
            // document.getElementById("dirs").innerHTML = ""
            // for (let i = 0; i < message.payload.dirs.length; i++) {
            //     index.addFolder(message.payload.dirs[i].name, message.payload.dirs[i].path);
            // }
            //
            // // Process files
            // document.getElementById("files_count").innerHTML = message.payload.files_count;
            // document.getElementById("files_size").innerHTML = message.payload.files_size;
            // document.getElementById("files").innerHTML = "";
            // if (typeof message.payload.files !== "undefined") {
            //     document.getElementById("files_panel").style.display = "block";
            //     let canvas = document.createElement("canvas");
            //     document.getElementById("files").append(canvas);
            //     new Chart(canvas, message.payload.files);
            // } else {
            //     document.getElementById("files_panel").style.display = "none";
            // }
        })
    },
    listen: function() {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    index.about(message.payload);
                    return {payload: "payload"};
                    break;
                case "check.out.menu":
                    asticode.notifier.info(message.payload);
                    break;
            }
        });
    }
};