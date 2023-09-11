import solid from "solid-start/vite";
import { defineConfig } from "vite";
import unocssPlugin from "unocss/vite";
export default defineConfig({
  plugins: [solid(), unocssPlugin()],
});
