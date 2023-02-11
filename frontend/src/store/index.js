import { createStore } from "redux";
import { taskReducer } from "./taskReducer";

export const store = createStore(taskReducer)

