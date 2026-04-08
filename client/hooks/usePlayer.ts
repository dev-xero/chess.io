import { PlayerContext } from '@/providers/player/Context';
import { useContext } from 'react';

export const usePlayer = () => {
    const context = useContext(PlayerContext);
    if (!context) {
        throw new Error('usePlayer must be used within a PlayerProvider');
    }
    return context;
};
