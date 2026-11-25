import { Test, TestingModule } from '@nestjs/testing';
import { AssetsGetAllService } from './assets.get_all.service';

describe('AssetsService', () => {
  let service: AssetsGetAllService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [AssetsGetAllService],
    }).compile();

    service = module.get<AssetsGetAllService>(AssetsGetAllService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
