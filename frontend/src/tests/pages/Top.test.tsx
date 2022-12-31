import React from 'react';
import { render, screen } from '@testing-library/react';
import Top from 'pages/Top';

test('renders learn react link', () => {
  render(<Top />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
