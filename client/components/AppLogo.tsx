import Image from "next/image";

export default function ChessIO() {
  return (
    <Image
      width={64}
      height={64}
      src="/logo.svg"
      alt="chess.io"
      className="mb-2 select-none"
      priority={true}
    />
  );
}
