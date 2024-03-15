import {
  Button,
  Card,
  CardActions,
  CardContent,
  Stack,
  Link,
  Typography,
  Skeleton,
  Alert,
} from '@mui/material';
import { getCurrentVersion } from 'api/version';
import { FC } from 'react';
import { useQuery } from 'react-query';
import { formatTimestamp } from 'utils/formatTimestamp';
import CachedIcon from '@mui/icons-material/Cached';
import { PMM_HOME_URL } from 'constants';
import { Messages } from './UpdateCard.messages';
export const UpdateCard: FC = () => {
  const { isLoading, data, error, isRefetching, refetch } = useQuery(
    ['currentVersion'],
    () => getCurrentVersion()
  );

  if (isLoading)
    return (
      <Card>
        <CardContent>
          <Stack spacing={1}>
            <Skeleton />
            <Skeleton />
          </Stack>
        </CardContent>
      </Card>
    );

  if (!data || error) {
    return (
      <Card>
        <CardContent>
          <Alert severity="error">{Messages.fetchError}</Alert>
        </CardContent>
      </Card>
    );
  }

  return (
    <Card sx={{ p: 1 }}>
      <CardContent>
        {data.installed.fullVersion === data.latest.fullVersion && (
          <Alert
            severity="success"
            sx={{
              mb: 2,
            }}
          >
            {Messages.upToDate}
          </Alert>
        )}
        <Stack spacing={1}>
          <Typography variant="body1">
            <Typography fontWeight="bold" component="strong">
              {Messages.runningVersion}
            </Typography>{' '}
            {data?.installed.version},{' '}
            {formatTimestamp(data?.installed.timestamp)}
          </Typography>
          <Typography variant="body1">
            <Typography fontWeight="bold" component="strong">
              {Messages.lastChecked}
            </Typography>{' '}
            {formatTimestamp(data?.lastCheck)}
          </Typography>
        </Stack>
      </CardContent>
      <CardActions>
        <Button
          startIcon={
            <CachedIcon
              sx={
                isRefetching
                  ? {
                      animation: 'spin 2s linear infinite',
                      '@keyframes spin': {
                        '0%': {
                          transform: 'rotate(360deg)',
                        },
                        '100%': {
                          transform: 'rotate(0deg)',
                        },
                      },
                    }
                  : {}
              }
            />
          }
          variant="contained"
          onClick={() => refetch()}
        >
          {isRefetching ? Messages.checking : Messages.checkNow}
        </Button>
        <Link href={PMM_HOME_URL}>
          <Button variant="outlined">{Messages.home}</Button>
        </Link>
      </CardActions>
    </Card>
  );
};
