// Generated by Wrangler by running `wrangler types --include-runtime=false` (hash: ba68d63ac044d23d69441491de7e0839)
declare namespace Cloudflare {
	interface Env {
		CONTAINER: DurableObjectNamespace<import("./src/index").Container>;
		CONTAINER_MANAGER: DurableObjectNamespace<import("./src/index").ContainerManager>;
	}
}
interface Env extends Cloudflare.Env {}
