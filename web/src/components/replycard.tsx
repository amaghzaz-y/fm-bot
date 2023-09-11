import { Answer } from "~/store";

export default function () {
  return (
    <>
      <div
        class="flex items-center justify-center 
      p-3 gap-2 bg-neutral-7 rounded-md max-w-xl text-justify"
      >
        {Answer() as string}
      </div>
    </>
  );
}
