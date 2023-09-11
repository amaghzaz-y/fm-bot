import { createSignal } from "solid-js";
import { createStore } from "solid-js/store";

export const [Answer, setAnswer] = createSignal<String>("");
