CREATE TABLE IF NOT EXISTS karaoke_sessions (
    id UUID PRIMARY KEY,
    nombre TEXT NOT NULL,
    hora_inicio TIMESTAMP NOT NULL,
    hora_fin TIMESTAMP NOT NULL,
    duracion_promedio_cancion INTEGER NOT NULL,
    estado TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS participants (
    id UUID PRIMARY KEY,
    session_id UUID REFERENCES karaoke_sessions(id) ON DELETE CASCADE,
    nombre TEXT NOT NULL,
    cancion TEXT NOT NULL,
    video_id TEXT,
    estado TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
