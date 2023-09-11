import { AiOutlineSend } from "solid-icons/ai";

export default function () {
  return (
    <>
      <div class="flex items-center justify-center p-2 gap-2">
        <input
          class="w-full max-w-xl outline-none border-none
          h-5 font-size-4 rounded-md p-3
          bg-neutral-9 text-neutral-1"
          placeholder="Enter a prompt here"
        />
        <button
          class="rounded-md outline-none h-11
          border-none bg-none bg-neutral-6
          hover:bg-neutral-9 text-center flex justify-center items-center"
        >
          <AiOutlineSend
            class="fill-white bg-none self-center place-self-center p-2 "
            size={"1.2rem"}
          />
        </button>
      </div>
    </>
  );
}
