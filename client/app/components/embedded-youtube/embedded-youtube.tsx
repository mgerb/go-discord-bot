import React from 'react';
import './embedded-youtube.scss';

interface IProps {
  id: string;
  className?: string;
}

export const EmbeddedYoutube = ({ id, className }: IProps) => {
  const src = `https://www.youtube.com/embed/${id}`;
  return (
    <div className={`embedded-youtube ${className}`}>
      <iframe src={src} className="embedded-youtube__iframe" allowFullScreen />
    </div>
  );
};
