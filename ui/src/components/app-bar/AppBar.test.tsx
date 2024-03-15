import { screen, render } from '@testing-library/react';
import { AppBar } from '.';
import { PMM_HOME_URL } from 'constants';
import { TestWrapper } from 'utils/testWrapper';

describe('AppBar', () => {
  it('links back to older PMM', () => {
    render(
      <TestWrapper>
        <AppBar />
      </TestWrapper>
    );

    expect(screen.getByTestId('appbar-pmm-link')).toHaveAttribute(
      'href',
      PMM_HOME_URL
    );
  });
});
