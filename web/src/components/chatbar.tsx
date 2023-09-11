import { AiOutlineSend } from "solid-icons/ai";
import { Answer, setAnswer } from "~/store";

export default function () {
  let input: any;

  return (
    <>
      <div class="flex items-center justify-center gap-2">
        <input
          ref={input}
          class="w-full max-w-xl outline-none border-none
          h-5 font-size-4 rounded-md p-3
          bg-neutral-9 text-neutral-1"
          placeholder="Enter a prompt here"
        />
        <button
          onClick={async () => {
            let res = await fetch(`http://127.0.0.1:1323/chat/${input.value}`);
            let answer = await res.text();
            console.log(answer);
            setAnswer(answer);
          }}
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
