import { copyFile, mkdir, readFile, writeFile } from "node:fs/promises";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const root = dirname(dirname(fileURLToPath(import.meta.url)));
const dist = join(root, "dist");
await mkdir(dist, { recursive: true });
await copyFile(join(root, "index.html"), join(dist, "index.html"));
await copyFile(join(root, "src", "style.css"), join(dist, "style.css"));
const source = await readFile(join(root, "src", "main.ts"), "utf8");
const js = source
  .replace(/type RebrandForm[\s\S]*?};\n\n/, "")
  .replace(/<RebrandForm>/g, "")
  .replace(/<HTMLPreElement>/g, "");
await writeFile(join(dist, "main.js"), js);
await writeFile(join(dist, "CNAME"), "rebrander.sh\n");

