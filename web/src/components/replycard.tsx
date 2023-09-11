import { Answer } from "~/store";

export default function () {
  return (
    <>
      <div
        class="flex items-center justify-center 
      p-2 gap-2 bg-neutral-7 rounded-md"
      >
        {Answer().toString()}
      </div>
    </>
  );
}
