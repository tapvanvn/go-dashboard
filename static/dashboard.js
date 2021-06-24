var _sock = null
var _connect_string = "ws://" + window.location.host+ "/ws"
var _items = {};
var _containers = {};

function getContainer(container_id) {

    if (typeof _containers[container_id] == 'undefined') {

        var el = document.createElement('div')
        el.className = "container"
        el.id = container_id
        el.innerHTML = "<span class='title'>" + container_id + "</span>"
        window.nct.core.binding("container", el)
    }
    return _containers[container_id];
}

function signal(item, params){

    var breadcums = item.split(".")
    var container_name = ""
    if (breadcums.length > 2) {
        console.log("nested container is not suppoted yet.", item)
        container_name = "default"
    }else {
        container_name = breadcums[0]
        item = breadcums[1]
    }
    var container = getContainer(container_name)
    var item = container.getItem(item)
    item.signal(params)
}

function connect() {

    if (_sock !== 'undefinned' && _sock != null && _sock.readyState == WebSocket.OPEN) {

        return
    }

    _sock = new WebSocket( _connect_string );

    _sock.onopen = function () {

        //on open
    }

    _sock.onclose = function (e) {

    }

    _sock.onmessage = function (e) {

        var msg = JSON.parse(e.data)

        console.log(msg)

        if (msg.item_name.length > 0) {
            
            if( typeof msg.params === 'object') {
                
                signal(msg.item_name, msg.params)
            } else {
                signal(msg.item_name, {})
            }
        }
    }
}

function Setup(){

    var Class = __pure__mod__.Class
	var p = __pure__mod__.Pure

    //MARK: container
    var container = __pure__mod__.Class.extend ( "Container",{

        init : function (dom, ctx) {
            this.dom = dom 
            this.ctx = ctx
            this.items = {}
            _containers[dom.id] = this
            p.dom.appendRoot(dom)
        },
        getItem: function(item_id){

            if (typeof this.items[item_id] == 'undefined') {
        
                var el = document.createElement('div')
                el.innerHTML = "<span class='title'>" + item_id + "</span>"
                el.className = "item"
                el.id = item_id
                handler = window.nct.core.binding("item", el)
                if (handler == null) {
                    console.log("binding item fail")
                }
                handler.setContainer(this)
                this.items[item_id] = handler
                this.dom.appendChild(el)
                
            }
            return this.items[item_id];
        },
        signal: function(params){
            
            var self = this 

            Object.keys(params).forEach((key)=>{
                
                var param = self.getParam(key) 
                $(param).find('.value').innerHTML(params[key])
            })
        }
    })

    window.nct.core.regType ( "container", container )

    //MARK: item
    var item = __pure__mod__.Class.extend ( "Item",{

        init : function (dom, ctx) {
            this.dom = dom 
            this.ctx = ctx
            this.params = {}
            _items[dom.id] = this
            this.lastSignal = Math.floor(Date.now() / 1000)
            this.container = null
            var self = this
            this.timer = setTimeout(() => {
                var now = Math.floor(Date.now() / 1000)
                if (now - self.lastSignal > 5) {
                    p.dom.unbindStyle(self.dom, 'active')
                }
            }, 5);
        },
        setContainer: function(container) {
            this.container = container
        },
        getParam: function (key) {
            var self = this 
            if (typeof self.params[key] != 'undefined') {
                return self.params[key]
            } 

            var el = document.createElement('div')
            el.className = "param"
            el.id = key
            el.innerHTML = "<span class='key'>" + key + "</span><span class='value'></span>"
            self.params[key] = el
            this.dom.appendChild(el)
            return el
        },
        signal: function(params){
            
            var self = this 

            Object.keys(params).forEach((key)=>{
                
                var param = self.getParam(key) 
                $(param).find('.value').innerHTML(params[key])
            })
            this.lastSignal = Math.floor(Date.now() / 1000)
            p.dom.bindStyle(this.dom, 'active')
        }
    })

    window.nct.core.regType ( "item", item )

    connect()
}

__pure__waiting__fn.push ( function (){
	var evt_handle = new __pure__mod__.EventHandle ({
		Context : this , fn : function ( evt ){
			
        Setup()
	}})
	if( typeof ( window.nct )!== 'undefined'){
		evt_handle.fn ()
	}
	else{
		window.__pure__.onTrigger ( "nct.init", evt_handle )
	}
})