using Priority_Queue;
using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace DiningPhilosophers
{
    class RequestQueue
    {
        private readonly SimplePriorityQueue<Request> _requests = new SimplePriorityQueue<Request>((current, next) => (int)(next - current));
        private readonly ConcurrentHashSet<string> _requestNames = new ConcurrentHashSet<string>();
        private readonly ConcurrentHashSet<string> _pendingRequests = new ConcurrentHashSet<string>();
        private Mutex _mux = new Mutex();

        private const int MaxRequestNames = 10;

        public RequestQueue()
        {
            EventManager.Subscribe("Finished", (name) => _pendingRequests.Remove(name), false);
            EventManager.Subscribe("RequestAdded", (name) => Run());
        }

        public void AddRequest(Philosopher philosopher)
        {
            _requestNames.Add(philosopher.Name);
            _requests.Enqueue(new Request { Philosopher = philosopher }, 0);
            EventManager.Broadcast("RequestAdded");
        }

        private void Run()
        {
            while (_requests.Count > 0)
            {
                _mux.WaitOne();
                var request = _requests.Dequeue();
                _mux.ReleaseMutex();
                var philosopher = request.Philosopher;

                var leftNeighborRequested = _requestNames.Contains(philosopher.LeftPhilosopher.Name);
                var rightNeighborRequested = _requestNames.Contains(philosopher.RightPhilosopher.Name);
                var lessThanTwoNeighborsRequested = !(leftNeighborRequested && rightNeighborRequested);
                var leftNeighborGranted = _pendingRequests.Contains(philosopher.LeftPhilosopher.Name);
                var rightNeighborGranted = _pendingRequests.Contains(philosopher.RightPhilosopher.Name);
                var neighborGranted = leftNeighborGranted || rightNeighborGranted;

                if (philosopher.CanEat && !neighborGranted && (request.Overdue || lessThanTwoNeighborsRequested || _requestNames.Length > MaxRequestNames))
                {
                    _requestNames.Remove(philosopher.Name);
                    _pendingRequests.Add(philosopher.Name);
                    EventManager.Broadcast(philosopher.Name + "RequestGranted");
                }
                else
                {
                    request.Priority++;
                    _mux.WaitOne();
                    _requests.Enqueue(request, request.Priority);
                    _mux.ReleaseMutex();
                }
            }
            EventManager.Subscribe("RequestAdded", (name) => Run());
        }
    }

    class Request
    {
        private const int MaxPriority = 30;

        public Philosopher Philosopher { get; set; }

        public int Priority { get; set; } = 0;

        public bool Overdue => Priority >= MaxPriority;
    }
}
