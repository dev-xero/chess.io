interface IErrorProps {
  err: string;
}

export default function Error(props: IErrorProps) {
  return props.err ? (
    <p className="text-center text-sm text-red-400 my-4 mb-8">{props.err}</p>
  ) : (
    <></>
  );
}
