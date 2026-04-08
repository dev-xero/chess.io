// 'use client';

import ProtectedPage from "@/components/ProtectedPage";
import HomeFragment from "@/fragments/HomeFragment";
import { PlayerProvider } from "@/providers/PlayerProvider";

export default function Home() {
  return (
    <ProtectedPage>
      <PlayerProvider>
        <HomeFragment />
      </PlayerProvider>
    </ProtectedPage>
  );
}
