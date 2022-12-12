import { useParams } from "react-router-dom";
import useGetNotes from "../../hooks/useGetNotes";

export default function Home() {
  const params = useParams();
  const { data } = useGetNotes();
  console.log(params);

  return (
    <div className="flex-auto">
      <div className="max-w-prose px-8 mx-auto">
        <h2 className="mt-32">Note Title</h2>
        <p className="mt-4">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Adipisci consequatur quo
          inventore recusandae magnam nemo esse eius consequuntur cumque ullam et nam odio eos
          facere numquam, exercitationem ipsa, a nulla?
        </p>
      </div>
    </div>
  );
}
