import "./lib";
import "./aspect";
import {entry} from "./lib";
import aspect from "./aspect";

export * from "./lib";

entry.buildAspect = () => {
    return new aspect()
}
