import "./lib/lib";
import "./aspect";
import { entry } from "./lib/lib";
import aspect from "./aspect";

export * from "./lib/lib";

entry.buildAspect = () => {
    return new aspect()
}
