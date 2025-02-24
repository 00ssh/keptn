import request from 'supertest';

describe('Test /bridgeInfo', () => {
  it('should return bridgeInfo', async () => {
    const response = await request(global.app).get('/api/bridgeInfo');
    expect(response.body).toEqual({
      bridgeVersion: 'develop',
      apiUrl: global.baseUrl,
      apiToken: 'apiToken',
      cliDownloadLink: 'https://github.com/keptn/keptn/releases',
      enableVersionCheckFeature: true,
      showApiToken: true,
      authType: 'NONE',
    });
    expect(response.statusCode).toBe(200);
  });
});
