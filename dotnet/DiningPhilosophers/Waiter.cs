using System;
using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;

namespace DiningPhilosophers
{
    class Waiter
    {
        private RequestQueue _requestQueue = new RequestQueue();

        public async Task<bool> Request(Philosopher philosopher)
        {
            var completionSource = new TaskCompletionSource<bool>();
            EventManager.Subscribe(philosopher.Name + "RequestGranted", (name) => 
            {
                completionSource.SetResult(true);
            });
            _requestQueue.AddRequest(philosopher);
            await completionSource.Task.ConfigureAwait(false);
           
            return true;   
        }
    }
}
