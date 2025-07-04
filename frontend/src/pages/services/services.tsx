import { useNavigate } from 'react-router';
import { useTranslation } from 'react-i18next';

import { Spinner } from '@/components/spinner';
import { ServicesTable } from '@/modules/services/services-table';
import { useServices } from '@/modules/services/hooks/use-services';
import { DataTableContainer } from '@/components/data-table-container';

function Services() {
  const navigate = useNavigate();
  const { t } = useTranslation('services');
  const { isLoading, data } = useServices();

  if (isLoading) {
    return (
      <div className='flex h-full w-full items-center justify-center'>
        <Spinner />
      </div>
    );
  }

  return (
    <DataTableContainer
      title={t('title')}
      description={t('description')}
      addNewButtonTitle={t('add')}
      onAddClick={() => navigate('/app/services/new')}
    >
      <div className='flex w-full flex-col gap-4'>
        <ServicesTable services={data ?? []} />
      </div>
    </DataTableContainer>
  );
}

export { Services };
