export namespace models {
	
	export class ConnectionConfig {
	    id: string;
	    name: string;
	    driver: string;
	    host: string;
	    port: number;
	    username: string;
	    password: string;
	    database: string;
	    databases: string[];
	
	    static createFrom(source: any = {}) {
	        return new ConnectionConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.driver = source["driver"];
	        this.host = source["host"];
	        this.port = source["port"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.database = source["database"];
	        this.databases = source["databases"];
	    }
	}

}

