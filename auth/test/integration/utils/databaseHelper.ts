/* eslint-disable @typescript-eslint/no-explicit-any */
// mongoose collection truncator

export const truncate = async (models: any) => {
  return await Promise.all(
    Object.keys(models).map(async (key) => {
      return await models[key].deleteMany({});
    }),
  );
};
