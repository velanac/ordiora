import { useMutation } from '@tanstack/react-query';

import { agent } from '@/lib/agent';
import { DoctorEventPayload } from '@/types';
import { ToastService } from '@/lib/toast-service';
import { queryClient, queryKeys } from '@/lib/query-client';

export const useCreateDoctorEvent = () =>
  useMutation({
    mutationFn: (data: DoctorEventPayload) =>
      agent.Events.createDoctorEvent(data),
    onSuccess: () => {
      ToastService.success('Doctor event created successfully');
      queryClient.invalidateQueries({
        queryKey: [queryKeys.events, 'offices'],
        type: 'all',
      });
    },
  });
